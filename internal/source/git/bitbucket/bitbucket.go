package bitbucket

import (
	"context"
	"devxstats/internal/model"
	"devxstats/internal/util"
	"fmt"
	"net/http"
)

type BitbucketProject struct {
	Description string `json:"description"`
	Namespace   string `json:"namespace"`
	Avatar      string `json:"avatar"`
	Scope       string `json:"scope"`
	Name        string `json:"name"`
	Key         string `json:"key"`
	ID          int    `json:"id"`
	Type        string `json:"type"`
	Public      bool   `json:"public"`
}

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
	fmt.Println("creating %v client, endpoint: ", system, baseURL)

	c := &client{
		baseURL:    fmt.Sprintf("%s%s", baseURL, apiPath),
		token:      token,
		httpClient: util.NewBearerHttpClient(token),
	}
	_, err := c.GetGroups(context.Background()) // use GetGroups as a test call
	if err != nil {
		return nil, fmt.Errorf("error creating %v client: %w", system, err)
	}

	return c, nil
}

func (c *client) Name() string {
	return system
}

func (c *client) GetGroups(ctx context.Context) ([]*model.Group, error) {
	fmt.Println("fetching groups")
	groups := []*model.Group{}
	// pagedGroups := &Page{Values: groups}

	// r, err := c.httpClient.Get(fmt.Sprintf("%s/projects", c.baseURL))
	// if err != nil {
	// 	fmt.Errorf("error fetching projects: %w", err)
	// }

	// defer r.Body.Close()
	// err = json.NewDecoder(r.Body).Decode(pagedGroups)
	// if err != nil {
	// 	return nil, fmt.Errorf("error decoding groups from api response: %w", err)
	// }

	return groups, nil
}

func (c *client) GetRepositories(ctx context.Context) ([]*model.Repo, error) {
	return []*model.Repo{{}}, nil
}

func (c *client) GetOpenPullRequests(ctx context.Context) ([]*model.PullRequest, error) {
	return []*model.PullRequest{{}}, nil
}
