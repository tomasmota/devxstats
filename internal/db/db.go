package db

import (
	"context"
	"devxstats/internal/model"
)

type DB interface {
	GetSystems(ctx context.Context) ([]*model.System, error)
	GetSystemByName(ctx context.Context, name string) (*model.System, error)
	AddGroup(context.Context, model.Group) error
	GetGroup(ctx context.Context, groupID int) (*model.Group, error)
	GetGroups(ctx context.Context) ([]*model.Group, error)
	GetGroupsBySystem(ctx context.Context, system model.System) ([]*model.Group, error)
	AddRepo(context.Context, model.Repo) error
	GetRepo(ctx context.Context, repoID int) (*model.Repo, error)
	GetRepos(ctx context.Context, groupID int) (*model.Repo, error)
}
