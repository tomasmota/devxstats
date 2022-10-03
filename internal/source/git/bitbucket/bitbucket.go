package bitbucket

import (
	"devxstats/internal/util"
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
	_, err := c.httpClient.Get(fmt.Sprintf("%s/projects", c.baseURL))
	if err != nil {
		return fmt.Errorf("error fetching projects: %w", err)
	}
	return nil
}
