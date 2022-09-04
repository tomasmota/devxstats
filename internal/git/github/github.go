package github

import (
	"devxstats/model"
	"fmt"
	"net/http"

	"github.com/drone/go-scm/scm"
	"github.com/drone/go-scm/scm/driver/github"
	"github.com/drone/go-scm/scm/transport"
)

type GithubConfig struct {
	Token   string
	BaseUrl string
}

type githubClient struct {
	Client *scm.Client
}

func NewClient(config *GithubConfig) (*githubClient, error) {
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
		Transport: &transport.PrivateToken{
			Token: config.Token,
		},
	}

	fmt.Println("api endpoint: ", c.BaseURL) //TODO: remove after testing

	return &githubClient{Client: c}, nil
}

func (githubClient) GetOpenPullRequests() ([]*model.PullRequest, error) {
	fmt.Println("Fetching open pull requests")
	prs := []*scm.PullRequest{{}} // TODO: fetch prs here
	return convertPullRequests(prs...), nil
}

func (githubClient) GetCommits() ([]*model.Commit, error) {
	fmt.Println("Fetching commits")
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
