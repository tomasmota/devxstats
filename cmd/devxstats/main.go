package main

import (
	"context"
	"devxstats/internal/config"
	"devxstats/internal/db"
	"devxstats/internal/server"
	"devxstats/internal/source/cd"
	"devxstats/internal/source/git"

	"fmt"
	"time"
)

// Fetches events from all sources and stores them on a loop
func syncSources(ctx context.Context, c config.AppConfig) {
	git := git.NewGitSyncer(c.Git)
	cd := cd.NewCdSyncer(c.Cd)
	// for {
	fmt.Println("\n---- Syncing Sources ----")
	err := git.Sync(ctx)
	if err != nil {
		panic(fmt.Errorf("error syncing git sources: %v", err))
	}
	err = cd.Sync(ctx)
	if err != nil {
		panic(fmt.Errorf("error syncing cd sources: %v", err))
	}
	time.Sleep(5000 * time.Millisecond)
	// }
}

func main() {
	ctx := context.Background()

	// Load configuration from environment variables and cmd args
	c := config.Load(ctx)

	// Initialize and ping database
	db.InitPostgres(ctx, c.Db)

	// Start syncing sources in the background
	go syncSources(ctx, *c)

	// Initialize http server
	app := &server.App{}
	app.InitializeRoutes()
	app.Run(fmt.Sprintf(":%d", c.Port))
}
