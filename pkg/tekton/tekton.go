package tekton

import "devxstats/model"

type Client interface {
	GetBuilds() ([]model.Build, error)
}

type clientImpl struct{}

func getClient() Client {
	return clientImpl{}
}

func (clientImpl) GetBuilds() ([]model.Build, error) {
	panic("unimplemented")
}
