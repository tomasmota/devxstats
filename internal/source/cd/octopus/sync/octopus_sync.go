package octopus_sync

import (
	"context"
	"devxstats/internal/db"
	"devxstats/internal/model"
	"devxstats/internal/util"
	"fmt"
	"net/url"

	od "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
)

const system = "octopus"

type OctopusSyncer struct {
	client *od.Client
	db     db.DB
}

func NewOctoupsSyncer(baseURL string, token string, db db.DB) (*OctopusSyncer, error) {
	fmt.Println("creating octopus syncer")
	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("an error occured while parsing octoups url: %w", err)
	}

	c, err := od.NewClient(util.NewHttpClient(), url, token, "")
	if err != nil {
		return nil, fmt.Errorf("error creating octopus client: %w", err)
	}

	syncer := &OctopusSyncer{
		client: c,
		db:     db,
	}
	return syncer, nil
}

func (s *OctopusSyncer) Sync(ctx context.Context) error {
	system, err := s.db.GetSystemByName(ctx, system)
	if err != nil {
		return fmt.Errorf("error fetching system: %w", err)
	}

	// Sync ProjectGroups
	projectGroups, err := s.client.ProjectGroups.GetAll()
	if err != nil {
		return fmt.Errorf("error fetching octopus ProjectGroups: %w", err)
	}
	for _, pg := range projectGroups {
		g := &model.Group{
			SystemID:    system.ID,
			Name:        pg.Name,
			Key:         pg.ID,
			Description: pg.Description,
		}
		s.db.AddGroup(ctx, *g)
	}

	// Sync Projects
	groups, err := s.db.GetGroupsBySystem(ctx, *system)
	if err != nil {
		return fmt.Errorf("error fetching octopus Projects: %w", err)
	}
	for _, g := range groups {
		p, err := s.client.ProjectGroups.GetByID(g.Key)
		if err != nil {
			return fmt.Errorf("error fetching octopus ProjectGroup %v: %w", g.Key, err)
		}
		projects, err := s.client.ProjectGroups.GetProjects(p)
		for _, p := range projects {
			pipeline := &model.CdPipeline{
				Name:    p.Name,
				GroupID: g.ID,
			}

			s.db.AddCdPipeline(ctx, *pipeline)
			if err != nil {
				return fmt.Errorf("error persisting octopus Projects: %w", err)
			}
		}
	}

	releases, err := s.client.Releases.Get()
	if err != nil {
		return fmt.Errorf("error fetching releases: %w", err)
	}
	for _, release := range releases.Items {
		deployResources, err := s.client.Deployments.GetDeployments(release)
		if err != nil {
			return fmt.Errorf("error fetching deployments in release %v: %w", release.ID, err)
		}
		odDeploys, err := deployResources.GetAllPages(s.client.Sling())
		if err != nil {
			return fmt.Errorf("error iterating through deployments in release %s: %w", release.ID, err)
		}
		for _, d := range odDeploys {
			deployment := &model.Deployment{
				PipelineId: 123, // fetch groupid
				StartedAt:  *d.Created,
				EndedAt:    *d.ModifiedOn, // TODO: map this properly
			}
			s.db.AddDeployment(ctx, *deployment)
		}
	}

	// Sync Deployments
	return nil
}
