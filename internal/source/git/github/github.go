package github

import (
	"context"
	"devxstats/internal/model"
	"devxstats/internal/util"
	"fmt"

	"github.com/drone/go-scm/scm"
	"github.com/drone/go-scm/scm/driver/github"
)

type githubClient struct {
	Client *scm.Client
}

const system = "github"

func NewClient(baseURL string, token string) (*githubClient, error) {
	fmt.Println("creating github client")
	var c *scm.Client
	var err error
	if baseURL != "" {
		c, err = github.New(baseURL)
		if err != nil {
			return nil, err
		}
	} else {
		c = github.NewDefault()
	}

	c.Client = util.NewBearerHttpClient(token)

	return &githubClient{Client: c}, nil
}

func (c *githubClient) Name() string {
	return system
}

func (c *githubClient) GetGroups(ctx context.Context) ([]*model.Group, error) {
	fmt.Println("get groups")
	orgs, res, err := c.Client.Organizations.List(ctx, scm.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("error fetching repositories: %w", err)
	}
	if res.Status != 200 {
		return nil, fmt.Errorf("error fetching repositories, received status: %v", res.Status)
	}
	for _, org := range orgs {
		fmt.Println(org.Name)
	}

	repos := []*scm.Organization{{}} // TODO: fetch prs here
	return convertGroups(repos...), nil
}

func convertGroups(from ...*scm.Organization) []*model.Group {
	// TODO: Implement
	return []*model.Group{{}}
}

func (c *githubClient) GetRepositories(ctx context.Context) ([]*model.Repo, error) {
	fmt.Println("fetching github repos")
	repos := []*scm.Repository{{}} // TODO: fetch prs here
	return convertRepositories(repos...), nil
}

func convertRepositories(from ...*scm.Repository) []*model.Repo {
	// TODO: Implement
	var to []*model.Repo
	for _, r := range from {
		to = append(to, &model.Repo{
			Name: r.Name,
		})
	}
	return to
}

func (c *githubClient) GetOpenPullRequests(ctx context.Context) ([]*model.PullRequest, error) {
	fmt.Println("fetching github open pull requests")
	prs := []*scm.PullRequest{{}} // TODO: fetch prs here
	return convertPullRequests(prs...), nil
}

func convertPullRequests(from ...*scm.PullRequest) []*model.PullRequest {
	// TODO: Implement
	return []*model.PullRequest{{}}
}
