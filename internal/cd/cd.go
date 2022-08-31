package cd

import (
	"devxstats/internal/cd/octopus"
	"devxstats/model"
)

type CdSyncer struct {
	sources []CdClient
}

type CdClient interface {
	GetDeployments() ([]*model.Deployment, error)
}

func NewCdSyncer( /*configuration of sources will somehow get injected into this method*/ ) *CdSyncer {
	syncer := &CdSyncer{}
	oc, err := octopus.NewOctopusClient("octopus.com")
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
