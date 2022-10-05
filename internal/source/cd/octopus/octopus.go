package octopus

import (
	"context"
	"devxstats/internal/model"
	"devxstats/internal/util"
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
	od "github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

const system = "octopus"

type client struct {
	client *od.Client
}

func (c *client) Name() string {
	return system
}

func NewClient(baseURL string, token string) (*client, error) {
	fmt.Println("creating octopus client")
	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("an error occured while parsing octoups url: %w", err)
	}

	c, err := od.NewClient(util.NewHttpClient(), url, token, "")
	if err != nil {
		return nil, fmt.Errorf("error creating octopus client: %w", err)
	}

	return &client{client: c}, nil
}

func (c *client) GetGroups(ctx context.Context) ([]*model.Group, error) {
	fmt.Printf("fetching %s groups", system)
	projectGroups, err := c.client.ProjectGroups.GetAll()
	if err != nil {
		return nil, fmt.Errorf("error fetching octopus ProjectGroups: %w", err)
	}

	var groups []*model.Group
	for _, pg := range projectGroups {
		groups = append(groups, &model.Group{
			Name:        pg.Name,
			Key:         pg.ID,
			Description: pg.Description,
		})
	}
	fmt.Printf("found %d groups", len(groups)) // TODO: remove after testing
	return groups, nil
}

func (c *client) GetCdPipelines(ctx context.Context, g model.Group) ([]*model.CdPipeline, error) {
	fmt.Printf("fetching %s pipelines", system)
	projects, err := c.client.Projects.GetAll()
	if err != nil {
		return nil, fmt.Errorf("error fetching octopus Projects: %w", err)
	}

	var pipelines []*model.CdPipeline
	for _, p := range projects { // TODO: find a solution that does not require getting all projects every time
		if p.ProjectGroupID == g.Key {
			pipelines = append(pipelines, &model.CdPipeline{
				Name:    p.Name,
				GroupID: g.ID,
			})
		}
	}

	fmt.Printf("found %d cd pipelines", len(pipelines)) // TODO: remove after testing
	return pipelines, nil
}

func (c *client) GetDeployments(ctx context.Context) ([]*model.Deployment, error) {
	fmt.Printf("fetching %s deployments", system)
	// iterate through releases and deployments within the releases

	d := []*octopusdeploy.Deployment{{}} // TODO: fetch deployments
	return convertDeployments(d...), nil
}

func convertDeployments(from ...*octopusdeploy.Deployment) []*model.Deployment {
	// TODO: Implement
	return []*model.Deployment{{}}
}
