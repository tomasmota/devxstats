package github

import (
	"context"
	"devxstats/internal/model"
	"devxstats/internal/util"
	"fmt"

	"github.com/google/go-github/v47/github"
)

type client struct {
	token  string
	client *github.Client
}

const system = "github"

func (c *client) Name() string {
	return system
}

func NewClient(token string) (*client, error) {
	fmt.Printf("creating %s client\n", system)

	c := &client{
		token:  token,
		client: github.NewClient(util.NewBearerHttpClient(token)),
	}

	err := c.ping()
	if err != nil {
		return nil, fmt.Errorf("error creating %v client: %w", system, err)
	}

	return c, nil
}

func (c *client) ping() error {
	_, _, err := c.client.RateLimits(context.Background())
	if err != nil {
		return fmt.Errorf("error testing connection to github: %w", err)
	}

	return nil
}

func (c *client) GetGroups(ctx context.Context) ([]*model.Group, error) {
	panic("unimplemented")
}

func (c *client) GetRepos(ctx context.Context) ([]*model.Repo, error) {
	panic("unimplemented")
}
