package tekton

import (
	"fmt"
)

type Deployment struct {
	// placeholder
}

var GetClient = getClient

type clientImpl struct {
	clusterUrl string
}

func getClient(clusterUrl string) clientImpl {
	return clientImpl{
		clusterUrl: clusterUrl,
	}
}

func (clientImpl) GetDeployments() ([]*Deployment, error) {
	fmt.Println("Fetching deployments from octopus")
	return []*Deployment{{}}, nil
}
