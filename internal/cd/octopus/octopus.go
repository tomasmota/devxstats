package octopus

import (
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
	_, err := url.Parse(config.BaseUrl)
	if err != nil {
		return nil, fmt.Errorf("an error occured while parsing octoups url: %v", err)
	}

	// c, err := od.NewClient(nil, url, config.Token, "")
	c := &od.Client{} // TODO: replace by real client once calling the real can load apikey from config
	if err != nil {
		log.Fatalf("error creating octopus client: %v", err)
	}

	return &octopusClient{Client: c}, nil
}

func (octopusClient) GetDeployments() ([]*model.Deployment, error) {
	fmt.Println("Fetching deployments")
	d := []*octopusdeploy.Deployment{{}} // TODO: Fetch commits here
	return convertDeployments(d...), nil
}

func convertDeployments(from ...*octopusdeploy.Deployment) []*model.Deployment {
	// TODO: Implement
	return []*model.Deployment{{}}
}
