package github

import (
	"devxstats/model"
	"fmt"

	"github.com/drone/go-scm/scm"
	"github.com/drone/go-scm/scm/driver/github"
)

type githubClient struct {
	Client *scm.Client
}

func NewGithubClient(baseUrl string) (*githubClient, error) {
	c, err := github.New(baseUrl)
	if err != nil {
		return nil, fmt.Errorf("an error occured while creating bitbucket client: %v", err)
	}

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
