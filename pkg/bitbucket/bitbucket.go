package bitbucket

import "devxstats/model"

type Client interface {
	GetOpenPullRequests() ([]model.PullRequest, error)
}

type clientImpl struct{}

func getClient() Client {
	return clientImpl{}
}

// GetOpenPullRequests implements Client
func (clientImpl) GetOpenPullRequests() ([]model.PullRequest, error) {
	panic("unimplemented")
}
