package db

import (
	"context"
	"devxstats/internal/config"
	"devxstats/internal/model"
	"fmt"
	"log"
	"time"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
)

// var pool *pgxpool.Pool

type pgdb struct {
	pool *pgxpool.Pool
}

func InitPostgres(ctx context.Context, c *config.DbConfig) DB {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	pool, err := pgxpool.Connect(ctx, "") // read from envs
	if err != nil {
		panic(fmt.Errorf("an error occured while creating database connection pool: %w", err))
	}

	err = pool.Ping(ctx)
	if err != nil {
		panic(fmt.Errorf("an error occured while pinging database: %w", err))
	}

	return &pgdb{pool: pool}
}

// GetSystems implements DB
func (db *pgdb) GetSystems(ctx context.Context) ([]*model.System, error) {
	var systems []*model.System
	err := pgxscan.Select(ctx, db.pool, &systems, `SELECT * FROM systems`)
	if err != nil {
		return nil, fmt.Errorf("error fetching systems: %w", err)
	}
	return systems, nil
}

// AddGroup implements store
func (db *pgdb) AddGroup(context.Context, model.Group) error {
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
func (db *pgdb) GetGroup(ctx context.Context, groupID int) (*model.Group, error) {
	panic("unimplemented")
}

// GetRepo implements store
func (db *pgdb) GetRepo(ctx context.Context, repoID int) (*model.Repo, error) {
	rows, err := db.pool.Query(ctx, "SELECT * FROM systems")
	if err != nil {
		return nil, fmt.Errorf("an error ocurred executing query: %w", err)
	}
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return nil, fmt.Errorf("an error while iterating through rows: %w", err)
		}

		for _, v := range values {
			log.Println(v)
		}
	}
	return nil, nil
}

// GetRepos implements store
func (db *pgdb) GetRepos(ctx context.Context, groupID int) (*model.Repo, error) {
	panic("unimplemented")
}
