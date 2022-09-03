package main

import (
	"context"
	"devxstats/internal/cd"
	"devxstats/internal/git"
	"devxstats/server"
	"devxstats/storage"
	"fmt"
	"log"
	"time"

	"github.com/sethvargo/go-envconfig"
)

type AppConfig struct {
	Port int `env:"PORT"`
}

func syncSources() {
	git := git.NewGitSyncer()
	cd := cd.NewCdSyncer()
	for {
		fmt.Println("\n---- Syncing Sources ----")
		git.Sync()
		cd.Sync()
		time.Sleep(1000 * time.Millisecond)
	}
}

func main() {
	ctx := context.Background()
	var c AppConfig
	if err := envconfig.Process(ctx, &c); err != nil {
		log.Fatal(err)
	}
	app := &server.App{}

	storage.InitializeDB()
	go syncSources()
	app.InitializeRoutes()
	app.Run(fmt.Sprintf(":%d", c.Port))
}
