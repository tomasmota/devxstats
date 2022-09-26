package main

import (
	"context"
	"devxstats/internal/api"
	"devxstats/internal/config"
	"devxstats/internal/db"
	"devxstats/internal/source/cd"
	"devxstats/internal/source/git"

	"fmt"
	"time"
)

// Fetches events from all sources and stores them on a loop
func syncSources(ctx context.Context, c config.AppConfig, db db.DB) {
	git := git.NewGitSyncer(c.Git, db)
	cd := cd.NewCdSyncer(c.Cd, db)
	// for {
	err := git.Sync(ctx)
	if err != nil {
		panic(fmt.Errorf("error syncing git sources: %w", err))
	}
	err = cd.Sync(ctx)
	if err != nil {
		panic(fmt.Errorf("error syncing cd sources: %w", err))
	}
	time.Sleep(5000 * time.Millisecond)
	// }
}

func main() {
	ctx := context.Background()

	// Load configuration from environment variables and cmd args
	c := config.Load(ctx)

	// Initialize and ping database
	db := db.InitPostgres(ctx, c.Db)

	// Start syncing sources in the background
	go syncSources(ctx, *c, db)

	// Initialize http server
	s := api.NewHTTPServer(db)
	s.Run(fmt.Sprintf(":%d", c.Port))
}
