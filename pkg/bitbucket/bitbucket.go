package bitbucket

import (
	"devxstats/model"
	"fmt"
)

// type User struct {
// 	Name         string `json:"name,omitempty"`
// 	EmailAddress string `json:"emailAddress,omitempty"`
// 	Id           int    `json:"id,omitempty"`
// 	DisplayName  string `json:"displayName,omitempty"`
// 	Active       bool   `json:"active,omitempty"`
// 	Slug         string `json:"slug,omitempty"`
// 	Type         string `json:"type,omitempty"`
// }

// type Author struct {
// 	User User `json:"user,omitempty"`
// }

// // https://developer.atlassian.com/server/bitbucket/rest/v803/api-group-projects/#api-projects-projectkey-repos-repositoryslug-pull-requests-get
// type PullRequest struct {
// 	Title       string   `json:"title,omitempty"`
// 	Description string   `json:"description,omitempty"`
// 	CreatedDate int64    `json:"createdDate,omitempty"` // Unix timestamp
// 	UpdatedDate int64    `json:"updatedDate,omitempty"` // Unix timestamp
// 	ClosedDate  int64    `json:"closedDate,omitempty"`  // Unix timestamp
// 	Status      string   `json:"status,omitempty"`
// 	Project     string   `json:"group,omitempty"` // Github org, Bitbucket project
// 	Repo        string   `json:"repo,omitempty"`
// 	Author      Author   `json:"author,omitempty"`
// 	Reviewers   []string `json:",omitempty"`
// }

// // https://developer.atlassian.com/server/bitbucket/rest/v803/api-group-projects/#api-projects-projectkey-repos-repositoryslug-commits-get
// type Commit struct {
// 	Timestamp time.Time `json:"timestamp,omitempty"`
// 	Team      string    `json:"team,omitempty"`
// 	Group     string    `json:"group,omitempty"` // Github org, Bitbucket project
// 	Repo      string    `json:"repo,omitempty"`
// 	User      string    `json:"user,omitempty"`
// }

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
