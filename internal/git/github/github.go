package github

import (
	"devxstats/model"
	"fmt"

	"github.com/drone/go-scm/scm"
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

func (clientImpl) GetOpenPullRequests() ([]*model.PullRequest, error) {
	fmt.Println("Fetching open pull requests")
	prs := []*scm.PullRequest{{}} // TODO: fetch prs here
	return convertPullRequests(prs...), nil
}

func (clientImpl) GetCommits() ([]*model.Commit, error) {
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
