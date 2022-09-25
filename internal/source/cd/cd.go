package cd

import (
	"context"
	"devxstats/internal/config"
	"devxstats/internal/db"
	"devxstats/internal/model"
	"devxstats/internal/source/cd/octopus"
)

type CdSyncer struct {
	sources []CdClient
	db      db.DB
}

type CdClient interface {
	GetDeployments(ctx context.Context) ([]*model.Deployment, error)
}

func NewCdSyncer(c *config.CdConfig, db db.DB) *CdSyncer {
	syncer := &CdSyncer{}
	if c.Octopus.Enabled {
		oc, err := octopus.NewOctopusClient(
			&octopus.OctopusConfig{
				BaseUrl: c.Octopus.Url,
				Token:   c.Octopus.Token,
			})
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
		_, err := source.GetDeployments(ctx)
		if err != nil {
			return err
		}

		// TODO: Persist Deployments
	}
	return nil
}
