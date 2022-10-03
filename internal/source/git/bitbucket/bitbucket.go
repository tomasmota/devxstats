package bitbucket

import (
	"context"
	"devxstats/internal/model"
	"devxstats/internal/util"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	apiPath = "/rest/api/1.0"
)

type client struct {
	baseURL    string
	token      string
	httpClient *http.Client
}

const system = "bitbucket"

func NewClient(baseURL string, token string) (*client, error) {
	fmt.Printf("creating %s client, endpoint: %s\n", system, baseURL)

	c := &client{
		baseURL:    fmt.Sprintf("%s%s", baseURL, apiPath),
		token:      token,
		httpClient: util.NewBearerHttpClient(token),
	}
	err := c.ping()
	if err != nil {
		return nil, fmt.Errorf("error creating %v client: %w", system, err)
	}

	return c, nil
}

func (c *client) Name() string {
	return system
}

func (c *client) ping() error {
	_, err := c.httpClient.Get(fmt.Sprintf("%s/projects", c.baseURL))
	if err != nil {
		return fmt.Errorf("error fetching projects: %w", err)
	}
	return nil
}

func (c *client) GetGroups(ctx context.Context) ([]*model.Group, error) {
	fmt.Println("fetching groups")
	var groups []*model.Group
	pagedGroups := &Page{}

	r, err := c.httpClient.Get(fmt.Sprintf("%s/projects", c.baseURL))
	if err != nil {
		return nil, fmt.Errorf("error fetching projects: %w", err)
	}

	defer r.Body.Close()

	err = json.NewDecoder(r.Body).Decode(pagedGroups)
	if err != nil {
		return nil, fmt.Errorf("error decoding groups from api response: %w", err)
	}

	for _, p := range pagedGroups.Projects {
		g := p.toGroup()
		groups = append(groups, g)
		fmt.Println(g.Name)
	}

	//TODO: do this on a loop
	return groups, nil
}

func (c *client) GetRepositories(ctx context.Context) ([]*model.Repo, error) {
	return []*model.Repo{{}}, nil
}

func (c *client) GetOpenPullRequests(ctx context.Context) ([]*model.PullRequest, error) {
	return []*model.PullRequest{{}}, nil
}
