package tekton

import (
	"devxstats/model"
	"fmt"
)

var GetClient = getClient

type clientImpl struct {
	clusterUrl string
}

func getClient(clusterUrl string) clientImpl {
	return clientImpl{
		clusterUrl: clusterUrl,
	}
}

func (clientImpl) GetBuilds() ([]model.Build, error) {
	fmt.Println("Fetching builds from tekton")
	return nil, nil
}
