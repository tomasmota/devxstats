package storage

import (
	"context"
	"devxstats/internal/config"
	"devxstats/internal/model"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type store interface {
	AddGroup(context.Context, model.Group) error
	GetGroup(groupID int) (model.Group, error)
	AddRepo(context.Context, model.Repo) error
	GetRepo(repoID int) (model.Repo, error)
	GetRepos(groupID int) (model.Repo, error)
}

type storeImpl struct {
	db *pgxpool.Pool
}

func InitializeDB(ctx context.Context, c *config.DbConfig) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, "") // read from envs
	if err != nil {
		panic(fmt.Errorf("an error occured while creating database connection pool: %v", err))
	}

	err = pool.Ping(ctx)
	if err != nil {
		panic(fmt.Errorf("an error occured while pinging database: %v", err))
	}

	initStore(&storeImpl{db: pool})
}

var DBStore store

func initStore(store store) {
	DBStore = store
}
