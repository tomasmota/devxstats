package github

import (
	"context"
	"devxstats/internal/model"
	"devxstats/internal/util"
	"fmt"

	"github.com/drone/go-scm/scm"
	"github.com/drone/go-scm/scm/driver/github"
)

type GithubConfig struct {
	BaseUrl string
	Token   string
}

type githubClient struct {
	Client *scm.Client
}

const system = "github"

// Creates a new client from *GithubConfig
func NewClient(config *GithubConfig) (*githubClient, error) {
	fmt.Println("creating github client")
	var c *scm.Client
	var err error
	if config.BaseUrl != "" {
		c, err = github.New(config.BaseUrl)
		if err != nil {
			return nil, err
		}
	} else {
		c = github.NewDefault()
	}

	c.Client = util.NewBearerHttpClient(config.Token)

	return &githubClient{Client: c}, nil
}

func (c *githubClient) Name() string {
	return system
}

func (c *githubClient) GetRepositories(ctx context.Context) ([]*model.Repo, error) {
	fmt.Println("fetching github open pull requests")
	repos := []*scm.Repository{{}} // TODO: fetch prs here
	return convertRepositories(repos...), nil
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
