package cd

import (
	"devxstats/internal/cd/octopus"
	"devxstats/internal/config"
	"devxstats/model"
)

type CdSyncer struct {
	sources []CdClient
}

type CdClient interface {
	GetDeployments() ([]*model.Deployment, error)
}

func NewCdSyncer(c *config.CdConfig) *CdSyncer {
	syncer := &CdSyncer{}
	oc, err := octopus.NewOctopusClient(
		&octopus.OctopusConfig{
			BaseUrl: c.Octopus.Url,
			Token:   c.Octopus.Token,
		})
	if err != nil {
		panic(err)
	}
	syncer.sources = append(syncer.sources, oc)

	return syncer
}

func (git *CdSyncer) Sync() error {
	for _, source := range git.sources {
		_, err := source.GetDeployments()
		if err != nil {
			return err
		}

		// TODO: Persist Deployments
	}
	return nil
}
