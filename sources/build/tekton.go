package sources

import (
	"devxstats/model"
	"fmt"
)

type TektonSource struct {
	clusterUrl string
	//...
}

func newTektonSource() *TektonSource {
	return &TektonSource{clusterUrl: "tekton.cluster.local"}
}

func (tektonSource *TektonSource) GetBuilds() ([]model.Build, error) {
	fmt.Println("Fetching builds from tekton")
	return nil, nil
}
