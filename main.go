package main

import (
	"context"
	"devxstats/internal/cd"
	"devxstats/internal/config"
	"devxstats/internal/git"
	"devxstats/server"
	"devxstats/storage"
	"fmt"
	"time"
)

// Fetches events from all sources and stores them on a loop
func syncSources(ctx context.Context, c config.AppConfig) {
	git := git.NewGitSyncer(c.Git)
	cd := cd.NewCdSyncer(c.Cd)
	for {
		fmt.Println("\n---- Syncing Sources ----")
		git.Sync(ctx)
		cd.Sync(ctx)
		time.Sleep(5000 * time.Millisecond)
	}
}

func main() {
	ctx := context.Background()

	// Load configuration from environment variables and cmd args
	c := config.Load(ctx)

	// Initialize and ping database
	storage.InitializeDB(ctx, c.Db)

	// Start syncing sources in the background
	go syncSources(ctx, *c)

	// Initialize http server
	app := &server.App{}
	app.InitializeRoutes()
	app.Run(fmt.Sprintf(":%d", c.Port))
}
