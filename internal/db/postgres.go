package db

import (
	"context"
	"devxstats/internal/model"
	"fmt"
	"time"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

type pgdb struct {
	pool *pgxpool.Pool
}

func InitPostgres(ctx context.Context) DB {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, "") // config gets read from envs: https://www.postgresql.org/docs/current/libpq-envars.html
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

// -------------------- SYSTEMS --------------------//
func (db *pgdb) GetSystems(ctx context.Context) ([]*model.System, error) {
	var systems []*model.System
	err := pgxscan.Select(ctx, db.pool, &systems, `SELECT * FROM systems`)
	if err != nil {
		return nil, fmt.Errorf("error fetching systems: %w", err)
	}
	return systems, nil
}

func (db *pgdb) GetSystemByName(ctx context.Context, name string) (*model.System, error) {
	var system model.System
	q := fmt.Sprintf(`SELECT * FROM systems WHERE name='%s'`, name)
	err := pgxscan.Get(ctx, db.pool, &system, q)
	if err != nil {
		return nil, fmt.Errorf("error fetching systems: %w", err)
	}
	return &system, nil
}

// -------------------- GROUPS --------------------//
func (db *pgdb) AddGroup(ctx context.Context, g model.Group) error {
	const sql = `
		INSERT INTO groups (system_id, name, description, key) 
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (system_id, name)
		DO
			UPDATE SET description = EXCLUDED.description;
	`
	_, err := db.pool.Exec(ctx, sql, g.SystemID, g.Name, g.Description, g.Key)
	if err != nil {
		return fmt.Errorf("error inserting group into database: %w", err)
	}
	return nil
}

func (db *pgdb) GetGroup(ctx context.Context, groupID int) (*model.Group, error) {
	panic("unimplemented")
}

func (db *pgdb) GetGroups(ctx context.Context) ([]*model.Group, error) {
	var groups []*model.Group
	err := pgxscan.Select(ctx, db.pool, &groups, `SELECT * FROM groups`)
	if err != nil {
		return nil, fmt.Errorf("error fetching groups: %w", err)
	}
	return groups, nil
}

func (db *pgdb) GetGroupsBySystem(ctx context.Context, system model.System) ([]*model.Group, error) {
	panic("unimplemented")
}

// -------------------- REPOS --------------------//
func (db *pgdb) AddRepo(ctx context.Context, r model.Repo) error {
	const sql = `
		INSERT INTO repos (group_id, scm_id, name, slug) 
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (group_id, scm_id)
		DO
			UPDATE SET group_id = EXCLUDED.group_id;
	`
	_, err := db.pool.Exec(ctx, sql, r.GroupID, r.ScmID, r.Name, r.Slug)
	if err != nil {
		return fmt.Errorf("error inserting repo into database: %w", err)
	}
	return nil
}

func (db *pgdb) GetRepo(ctx context.Context, repoID int) (*model.Repo, error) {
	panic("unimplemented")
}

func (db *pgdb) GetRepos(ctx context.Context, groupID int) (*model.Repo, error) {
	panic("unimplemented")
}

func (*pgdb) AddCdPipeline(context.Context, model.CdPipeline) error {
	panic("unimplemented")
}

func (*pgdb) AddDeployment(context.Context, model.Deployment) error {
	panic("unimplemented")
}
