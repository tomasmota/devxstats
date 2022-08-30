package main

import (
	"devxstats/server"
	build "devxstats/sources/build"
	git "devxstats/sources/git"
	"devxstats/storage"
	"fmt"
	"time"
)

func syncSources() {
	gitSources := git.NewGitSources()
	buildSources := build.NewBuildSources()
	for {
		fmt.Println("Syncing Sources")
		gitSources.Sync()
		buildSources.Sync()
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
