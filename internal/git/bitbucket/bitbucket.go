package bitbucket

import (
	"devxstats/model"
	"fmt"

	"github.com/drone/go-scm/scm"
	"github.com/drone/go-scm/scm/driver/stash"
)

type bitbucketClient struct {
	Client *scm.Client
}

func NewBitbucketClient(baseUrl string) (*bitbucketClient, error) {
	c, err := stash.New(baseUrl)
	if err != nil {
		return nil, fmt.Errorf("an error occured while creating bitbucket client: %v", err)
	}

	return &bitbucketClient{Client: c}, nil
}

func (c *bitbucketClient) GetOpenPullRequests() ([]*model.PullRequest, error) {
	fmt.Println("Fetching open pull requests")
	prs := []*scm.PullRequest{{}} // TODO: fetch prs here
	return convertPullRequests(prs...), nil
}

func (c *bitbucketClient) GetCommits() ([]*model.Commit, error) {
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
