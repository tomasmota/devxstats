package main

import (
	"devxstats/internal/cd"
	"devxstats/internal/git"
	"devxstats/server"
	"devxstats/storage"
	"fmt"
	"time"
)

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
	app := &server.App{}

	storage.InitializeDB()
	go syncSources()
	app.InitializeRoutes()
	app.Run(":8080")
}
