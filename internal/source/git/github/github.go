package github

import (
	"context"
	"devxstats/internal/model"
	"devxstats/internal/util"
	"fmt"
	"net/http"
)

const (
	apiPath = "/something/here" // TODO: set this proper
)

type client struct {
	baseURL    string
	token      string
	httpClient *http.Client
}

const system = "github"

func (c *client) Name() string {
	return system
}

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

func (c *client) ping() error {
	// TODO: call some random endpoint just for testing connection
	return nil
}

func (c *client) GetGroups(ctx context.Context) ([]*model.Group, error) {
	panic("unimplemented")
}

func (c *client) GetRepos(ctx context.Context) ([]*model.Repo, error) {
	panic("unimplemented")
}
