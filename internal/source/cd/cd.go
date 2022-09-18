package cd

import (
	"context"
	"devxstats/internal/config"
	"devxstats/internal/model"
	"devxstats/internal/source/cd/octopus"
)

type CdSyncer struct {
	sources []CdClient
}

type CdClient interface {
	GetDeployments(ctx context.Context) ([]*model.Deployment, error)
}

func NewCdSyncer(c *config.CdConfig) *CdSyncer {
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

	return syncer
}

func (git *CdSyncer) Sync(ctx context.Context) error {
	for _, source := range git.sources {
		_, err := source.GetDeployments(ctx)
		if err != nil {
			return err
		}

		// TODO: Persist Deployments
	}
	return nil
}
