package db

import (
	"context"
	"devxstats/internal/model"
	"fmt"
	"time"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

type pgdb struct {
	pool *pgxpool.Pool
}

func InitPostgres(ctx context.Context) DB {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	pool, err := pgxpool.Connect(ctx, "") // config gets read from envs: https://www.postgresql.org/docs/current/libpq-envars.html
	if err != nil {
		panic(fmt.Errorf("an error occured while creating database connection pool: %w", err))
	}

	err = pool.Ping(ctx)
	if err != nil {
		panic(fmt.Errorf("an error occured while pinging database: %w", err))
	}

	fmt.Println("db connection innitialized")
	return &pgdb{pool: pool}
}

func (db *pgdb) GetSystems(ctx context.Context) ([]*model.System, error) {
	var systems []*model.System
	err := pgxscan.Select(ctx, db.pool, &systems, `SELECT * FROM systems`)
	if err != nil {
		return nil, fmt.Errorf("error fetching systems: %w", err)
	}
	return systems, nil
}

func (db *pgdb) AddGroup(context.Context, model.Group) error {
	panic("unimplemented")
}

func (db *pgdb) AddRepo(ctx context.Context, repo model.Repo) error {
	panic("unimplemented")
}

func (db *pgdb) GetGroup(ctx context.Context, groupID int) (*model.Group, error) {
	panic("unimplemented")
}

func (db *pgdb) GetRepo(ctx context.Context, repoID int) (*model.Repo, error) {
	panic("unimplemented")
}

func (db *pgdb) GetRepos(ctx context.Context, groupID int) (*model.Repo, error) {
	panic("unimplemented")
}
