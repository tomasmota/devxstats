package sources

import (
	"devxstats/model"
	"devxstats/pkg/github"
	"devxstats/storage"
	"fmt"
	"time"
)

type GithubSource struct {
	baseUrl string
	client  github.Client
	//...
}

func newGithubSource() *GithubSource {
	return &GithubSource{baseUrl: "github.com"}
}

func (githubSource *GithubSource) GetCommits() ([]model.Commit, error) {
	fmt.Println("Fetching commits from github")
	events := []interface{}{
		model.Commit{
			Timestamp: time.Now(),
			Team:      "devx",
			Group:     "ABCD",
			Repo:      "devxstats",
			System:    "Github",
			User:      "tomas-mota",
		},
		model.Commit{
			Timestamp: time.Now(),
			Team:      "devx",
			Group:     "EFGH",
			Repo:      "devxstats2",
			System:    "bitbucket",
			User:      "jane doe",
		},
	}
	fmt.Printf("Adding %v commits \n", len(events))

	err := storage.DBStore.AddCommits(events)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (githubSource *GithubSource) GetOpenPullRequests() ([]model.PullRequest, error) {
	fmt.Println("Fetching open pull requests from github")
	openPullRequests, err := githubSource.client.GetOpenPullRequests()
	if err != nil {
		return nil, err
	}
	return openPullRequests, nil
}
