package sources

import (
	"devxstats/model"
	"devxstats/pkg/tekton"
	"fmt"
)

type TektonSource struct {
	clusterUrl string
	client     tekton.Client
	//...
}

func newTektonSource() *TektonSource {
	return &TektonSource{clusterUrl: "tekton.cluster.local"}
}

func (tektonSource *TektonSource) GetBuilds() ([]model.Build, error) {
	fmt.Println("Fetching builds from tekton")
	builds, err := tektonSource.client.GetBuilds()
	if err != nil {
		return nil, err
	}
	return builds, nil
}
