package octopus

import (
	"context"
	"devxstats/internal/util"
	"devxstats/model"
	"fmt"
	"log"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	od "github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

type OctopusConfig struct {
	BaseUrl string
	Token   string
}

type octopusClient struct {
	Client *od.Client
}

func NewOctopusClient(config *OctopusConfig) (*octopusClient, error) {
	fmt.Println("creating octopus client")
	url, err := url.Parse(config.BaseUrl)
	if err != nil {
		return nil, fmt.Errorf("an error occured while parsing octoups url: %v", err)
	}

	c, err := od.NewClient(util.NewHttpClient(), url, config.Token, "")
	if err != nil {
		log.Fatalf("error creating octopus client: %v", err)
	}

	return &octopusClient{Client: c}, nil
}

func (octopusClient) GetDeployments(ctx context.Context) ([]*model.Deployment, error) {
	fmt.Println("fetching deployments")
	d := []*octopusdeploy.Deployment{{}} // TODO: Fetch commits here
	return convertDeployments(d...), nil
}

func convertDeployments(from ...*octopusdeploy.Deployment) []*model.Deployment {
	// TODO: Implement
	return []*model.Deployment{{}}
}
