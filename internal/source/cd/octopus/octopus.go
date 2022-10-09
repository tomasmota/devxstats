package octopus

import (
	"context"
	"devxstats/internal/model"
	"devxstats/internal/util"
	"fmt"
	"net/url"

	od "github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/client"
)

const system = "octopus"

type odClient struct {
	client *od.Client
}

func (c *odClient) Name() string {
	return system
}

func NewClient(baseURL string, token string) (*odClient, error) {
	fmt.Println("creating octopus client")
	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("an error occured while parsing octoups url: %w", err)
	}

	c, err := od.NewClient(util.NewHttpClient(), url, token, "")
	if err != nil {
		return nil, fmt.Errorf("error creating octopus client: %w", err)
	}

	return &odClient{client: c}, nil
}

func (c *odClient) GetGroups(ctx context.Context) ([]*model.Group, error) {
	fmt.Printf("fetching %s groups\n", system)
	projectGroups, err := c.client.ProjectGroups.GetAll()
	if err != nil {
		return nil, fmt.Errorf("error fetching octopus ProjectGroups: %w", err)
	}
	var groups []*model.Group
	for _, pg := range projectGroups {
		groups = append(groups, &model.Group{
			Name:        pg.Name,
			Key:         pg.ID,
			Description: pg.Description,
		})
	}
	return groups, nil
}

func (c *odClient) GetCdPipelines(ctx context.Context, g model.Group) ([]*model.CdPipeline, error) {
	fmt.Printf("fetching %s pipelines\n", system)

	p, err := c.client.ProjectGroups.GetByID(g.Key)
	if err != nil {
		return nil, fmt.Errorf("error fetching octopus Project Group: %w", err)
	}

	projects, err := c.client.ProjectGroups.GetProjects(p)
	if err != nil {
		return nil, fmt.Errorf("error fetching octopus Projects: %w", err)
	}

	var pipelines []*model.CdPipeline
	for _, p := range projects {
		pipelines = append(pipelines, &model.CdPipeline{
			Name:    p.Name,
			GroupID: g.ID,
		})
	}

	fmt.Printf("found %d cd pipelines", len(pipelines)) // TODO: remove after testing
	return pipelines, nil
}

func (c *odClient) GetDeployments(ctx context.Context) ([]*model.Deployment, error) {
	// fmt.Printf("fetching %s deployments\n", system)

	// var deploys []*model.Deployment
	// releases, err := c.client.Releases.Get()
	// if err != nil {
	// 	return nil, fmt.Errorf("error fetching releases: %w", err)
	// }
	// for _, release := range releases.Items {
	// 	deployResources, err := c.client.Deployments.GetDeployments(release)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("error fetching deployments in release %v: %w", release.ID, err)
	// 	}
	// 	odDeploys, err := deployResources.GetAllPages(c.client.Sling())
	// 	if err != nil {
	// 		return nil, fmt.Errorf("error iterating through deployments: %w", release.ID, err)
	// 	}
	// 	for _, d := range odDeploys {
	// 		deploys = append(deploys, &model.Deployment{})
	// 	}

	// }
	return nil, nil
}
