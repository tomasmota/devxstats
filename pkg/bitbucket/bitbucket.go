package bitbucket

import (
	"devxstats/model"
	"fmt"
)

var GetClient = getClient

type clientImpl struct {
	baseUrl string
}

func getClient(baseUrl string) clientImpl {
	return clientImpl{
		baseUrl: baseUrl,
	}
}

func (clientImpl) GetOpenPullRequests() ([]model.PullRequest, error) {
	fmt.Println("Fetching open pull requests from bitbucket")
	return nil, nil
}

func (clientImpl) GetCommits() ([]model.Commit, error) {
	fmt.Println("Fetching commits from bitbucket")
	return nil, nil
}
