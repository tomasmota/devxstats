package sources

import (
	"devxstats/model"
	"fmt"
)

type GithubSource struct {
	baseUrl string
	//...
}

func newGithubSource() *GithubSource {
	return &GithubSource{baseUrl: "github.com"}
}

func (githubSource *GithubSource) GetCommits() ([]model.Commit, error) {
	fmt.Println("Fetching commits from github")
	return nil, nil
}

func (githubSource *GithubSource) GetOpenPullRequests() ([]model.PullRequest, error) {
	fmt.Println("Fetching open pull requests from github")
	return nil, nil
}
