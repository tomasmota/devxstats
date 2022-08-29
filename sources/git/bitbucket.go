package sources

import "devxstats/model"

type BitbucketSource struct {
	baseUrl string
	//...
}

func newBitbucketSource() *BitbucketSource {
	return &BitbucketSource{baseUrl: "changeme.com"}
}

func (bitbucketSource *BitbucketSource) GetCommits() ([]model.Commit, error) {
	return nil, nil
}

func (bitbucketSource *BitbucketSource) GetOpenPullRequests() ([]model.Commit, error) {
	return nil, nil
}
