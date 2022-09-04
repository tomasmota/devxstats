package main

import (
	"context"
	"devxstats/internal/cd"
	"devxstats/internal/config"
	"devxstats/internal/git"
	"devxstats/server"
	"devxstats/storage"
	"fmt"
	"log"
	"time"

	"github.com/sethvargo/go-envconfig"
)

// Fetches events from all sources and stores them on a loop
func syncSources(c config.AppConfig) {
	git := git.NewGitSyncer(c.Git)
	cd := cd.NewCdSyncer(c.Cd)
	for {
		fmt.Println("\n---- Syncing Sources ----")
		git.Sync()
		cd.Sync()
		time.Sleep(5000 * time.Millisecond)
	}
}

func main() {
	ctx := context.Background()
	var c config.AppConfig
	if err := envconfig.Process(ctx, &c); err != nil {
		log.Fatal(err)
	}
	app := &server.App{}

	storage.InitializeDB(ctx, c.Db)
	go syncSources(c)
	app.InitializeRoutes()
	app.Run(fmt.Sprintf(":%d", c.Port))
}
