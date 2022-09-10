package github

import (
	"context"
	"devxstats/model"
	"fmt"
	"net/http"

	"github.com/drone/go-scm/scm"
	"github.com/drone/go-scm/scm/driver/github"
	"github.com/drone/go-scm/scm/transport"
)

type GithubConfig struct {
	BaseUrl string
	Token   string
}

type githubClient struct {
	Client *scm.Client
}

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

	c.Client = &http.Client{
		Transport: &transport.BearerToken{
			Token: config.Token,
		},
	}

	return &githubClient{Client: c}, nil
}

func (c *githubClient) GetOpenPullRequests(ctx context.Context) ([]*model.PullRequest, error) {
	fmt.Println("Fetching github open pull requests")
	prs := []*scm.PullRequest{{}} // TODO: fetch prs here
	return convertPullRequests(prs...), nil
}

func (c *githubClient) GetCommits(ctx context.Context) ([]*model.Commit, error) {
	fmt.Println("Fetching github commits")

	commits := []*scm.Commit{{}} // TODO: Fetch commits here
	return convertCommits(commits...), nil
}

func convertPullRequests(from ...*scm.PullRequest) []*model.PullRequest {
	// TODO: Implement
	return []*model.PullRequest{{}}
}

func convertCommits(from ...*scm.Commit) []*model.Commit {
	// TODO: Implement
	return []*model.Commit{{}}
}
