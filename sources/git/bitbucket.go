package sources

import (
	"devxstats/model"
	"fmt"
)

type BitbucketSource struct {
	baseUrl string
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
	fmt.Println("Fetching open pull requests from bitbucket")
	return nil, nil
}
