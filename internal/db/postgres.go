package db

import (
	"context"
	"devxstats/internal/config"
	"devxstats/internal/model"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// var pool *pgxpool.Pool

type pgdb struct {
	pool *pgxpool.Pool
}

func InitPostgres(ctx context.Context, c *config.DbConfig) DB {
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

	return &pgdb{pool: pool}
}

// AddGroup implements store
func (*pgdb) AddGroup(context.Context, model.Group) error {
	panic("unimplemented")
}

// AddRepo implements store
func (db *pgdb) AddRepo(ctx context.Context, repo model.Repo) error {
	conn, err := db.pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("unable to acquire db conn %w", err)
	}
	_, err = conn.Exec(ctx, "SELECT * FROM systems") // TODO: add real query
	if err != nil {
		return fmt.Errorf("error running query %w", err)
	}
	return nil
}

// GetGroup implements store
func (*pgdb) GetGroup(int) (model.Group, error) {
	panic("unimplemented")
}

// GetRepo implements store
func (*pgdb) GetRepo(int) (model.Repo, error) {
	panic("unimplemented")
}

// GetRepos implements store
func (*pgdb) GetRepos(groupID int) (model.Repo, error) {
	panic("unimplemented")
}
