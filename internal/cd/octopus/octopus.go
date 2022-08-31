package octopus

import (
	"devxstats/model"
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

var GetClient = getClient

type clientImpl struct {
	baseUrl string
}

func getClient(baseUrl string) clientImpl {
	return clientImpl{
		baseUrl: baseUrl,
	}
}

func (clientImpl) GetDeployments() ([]*model.Deployment, error) {
	fmt.Println("Fetching deployments")
	d := []*octopusdeploy.Deployment{{}} // TODO: Fetch commits here
	return convertDeployments(d...), nil
}

func convertDeployments(from ...*octopusdeploy.Deployment) []*model.Deployment {
	// TODO: Implement
	return []*model.Deployment{{}}
}
