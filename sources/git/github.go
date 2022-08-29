package sources

import "devxstats/model"

type GithubSource struct {
	baseUrl string
	//...
}

func newGithubSource() *GithubSource {
	return &GithubSource{baseUrl: "changeme.com"}
}

func (githubSource *GithubSource) GetCommits() ([]model.Commit, error) {
	return nil, nil
}

func (githubSource *GithubSource) GetOpenPullRequests() ([]model.Commit, error) {
	return nil, nil
}
