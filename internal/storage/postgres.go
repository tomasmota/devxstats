package storage

import (
	"context"
	"devxstats/internal/model"
)

// AddGroup implements store
func (*storeImpl) AddGroup(context.Context, model.Group) error {
	panic("unimplemented")
}

// AddRepo implements store
func (*storeImpl) AddRepo(context.Context, model.Repo) error {
	panic("unimplemented")
}

// GetGroup implements store
func (*storeImpl) GetGroup(int) (model.Group, error) {
	panic("unimplemented")
}

// GetRepo implements store
func (*storeImpl) GetRepo(int) (model.Repo, error) {
	panic("unimplemented")
}

// GetRepos implements store
func (*storeImpl) GetRepos(groupID int) (model.Repo, error) {
	panic("unimplemented")
}
