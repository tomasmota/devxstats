package cd

import (
	"context"
	"devxstats/internal/config"
	"devxstats/internal/db"
	"devxstats/internal/model"
	"devxstats/internal/source/cd/octopus"
	"fmt"
)

type CdSyncer struct {
	sources []CdClient
	db      db.DB
}

type CdClient interface {
	GetGroups(ctx context.Context) ([]*model.Group, error)
	GetCdPipelines(ctx context.Context, group model.Group) ([]*model.CdPipeline, error)
	GetDeployments(ctx context.Context) ([]*model.Deployment, error)
	Name() string
}

func NewCdSyncer(c *config.CdConfig, db db.DB) *CdSyncer {
	syncer := &CdSyncer{}
	if c.Octopus.Enabled {
		oc, err := octopus.NewClient(c.Octopus.Url, c.Octopus.Token)
		if err != nil {
			panic(err)
		}
		syncer.sources = append(syncer.sources, oc)
	}

	syncer.db = db

	return syncer
}

func (s *CdSyncer) Sync(ctx context.Context) error {
	for _, source := range s.sources {
		system, err := s.db.GetSystemByName(ctx, source.Name())
		if err != nil {
			return fmt.Errorf("error fetching system entity for %s: %w", source.Name(), err)
		}

		// GROUPS
		groups, err := source.GetGroups(ctx)
		if err != nil {
			return err
		}

		for _, g := range groups {
			g.SystemID = system.ID
			err := s.db.AddGroup(ctx, *g)
			if err != nil {
				return fmt.Errorf("error persisting group %v: %w", g.Name, err)
			}

			fmt.Printf("%v: finished syncing groups\n", source.Name())
		}

		// CD PIPELINES
		// fetch groups again from database, before fetching pipelines within each
		groups, err = s.db.GetGroupsBySystem(ctx, *system)
		if err != nil {
			return fmt.Errorf("error fetching system entity for %s: %w", source.Name(), err)
		}
		for _, g := range groups {
			_, err = source.GetCdPipelines(ctx, *g)
			if err != nil {
				return err
			}
		}

		_, err = source.GetDeployments(ctx)
		if err != nil {
			return err
		}

		// TODO: Persist Deployments
	}
	return nil
}
