package sources

import (
	"devxstats/model"
	"devxstats/pkg/bitbucket"
	"fmt"
)

type BitbucketSource struct {
	baseUrl string
	client  bitbucket.Client
	//...
}

func newBitbucketSource() *BitbucketSource {
	return &BitbucketSource{baseUrl: "bitbucket.com"}
}

func (bitbucketSource *BitbucketSource) GetCommits() ([]model.Commit, error) {
	fmt.Println("Fetching commits from bitbucket")
	return nil, nil
}

func (bitbucketSource *BitbucketSource) GetOpenPullRequests() ([]model.PullRequest, error) {
	// https://{baseurl}/rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/pull-requests
	fmt.Println("Fetching open pull requests from bitbucket")
	openPullRequests, err := bitbucketSource.client.GetOpenPullRequests()
	if err != nil {
		return nil, err
	}
	return openPullRequests, nil
}
