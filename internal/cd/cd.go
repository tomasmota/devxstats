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
	git := &CdSyncer{}
	// Add sources based on configuration
	git.sources = append(git.sources, octopus.GetClient("octopus.com")) //TODO: pass in config params
	return git
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
