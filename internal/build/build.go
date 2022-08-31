package build

import (
	"devxstats/model"
	"devxstats/pkg/tekton"
)

type buildSources struct {
	sources []BuildSources
}

type BuildSources interface {
	GetBuilds() ([]model.Build, error)
}

func NewBuildSources( /*configuration of sources will somehow get injected into this method*/ ) *buildSources {
	buildSources := &buildSources{}
	// if config.contains("tekton") {
	buildSources.sources = append(buildSources.sources, tekton.GetClient("k8s.cluster.local"))
	// }
	return buildSources
}

func (buildSources *buildSources) Sync() error {
	for _, source := range buildSources.sources {
		_, err := source.GetBuilds()
		if err != nil {
			return err
		}

		//TODO: Persist to database
	}
	return nil
}
