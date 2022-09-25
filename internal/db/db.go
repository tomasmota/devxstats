package db

import (
	"context"
	"devxstats/internal/model"
)

type DB interface {
	AddGroup(context.Context, model.Group) error
	GetGroup(ctx context.Context, groupID int) (*model.Group, error)
	AddRepo(context.Context, model.Repo) error
	GetRepo(ctx context.Context, repoID int) (*model.Repo, error)
	GetRepos(ctx context.Context, groupID int) (*model.Repo, error)
}
