package storage

import (
	"context"
	"devxstats/internal/config"
	"devxstats/internal/model"

	"github.com/jackc/pgx/v5"
)

type store interface {
	AddGroup(context.Context, model.Group) error
	GetGroup(groupID int) (model.Group, error)
	AddRepo(context.Context, model.Repo) error
	GetRepo(repoID int) (model.Repo, error)
	GetRepos(groupID int) (model.Repo, error)
}

type storeImpl struct {
	db *pgx.Conn
}

func InitializeDB(ctx context.Context, c *config.DbConfig) {
	// ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	// defer cancel()

	// conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	// if err != nil {
	// 	panic(fmt.Errorf("an error occured while creating database connection: %v", err))
	// }

	conn := &pgx.Conn{}
	initStore(&storeImpl{db: conn})
}

var DBStore store

func initStore(store store) {
	DBStore = store
}
