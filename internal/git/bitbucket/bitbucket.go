package bitbucket

import (
	"context"
	"devxstats/model"
	"fmt"
	"net/http"

	"github.com/drone/go-scm/scm"
	"github.com/drone/go-scm/scm/driver/stash"
	"github.com/drone/go-scm/scm/transport"
)

type BitbucketConfig struct {
	BaseUrl string
	Token   string
}

type bitbucketClient struct {
	Client *scm.Client
}

func NewBitbucketClient(config *BitbucketConfig) (*bitbucketClient, error) {
	fmt.Println("creating bitbucket client, endpoint: ", config.BaseUrl)

	c, err := stash.New(config.BaseUrl)
	if err != nil {
		return nil, fmt.Errorf("an error occured while creating bitbucket client: %v", err)
	}

	c.Client = &http.Client{
		Transport: &transport.PrivateToken{
			Token: config.Token,
		},
	}

	return &bitbucketClient{Client: c}, nil
}

func (c *bitbucketClient) GetOpenPullRequests(ctx context.Context) ([]*model.PullRequest, error) {
	fmt.Println("Fetching bitbucket open pull requests")
	prs := []*scm.PullRequest{{}} // TODO: fetch prs here
	return convertPullRequests(prs...), nil
}

func (c *bitbucketClient) GetCommits(ctx context.Context) ([]*model.Commit, error) {
	fmt.Println("Fetching bitbucket commits")
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
