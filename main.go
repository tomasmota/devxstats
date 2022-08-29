package main

import (
	"devxstats/server"
	sources "devxstats/sources/git"
	"devxstats/storage"
	"fmt"
	"time"
)

func syncSources() {
	gitSource := sources.NewGit()
	for {
		fmt.Println("Syncing Sources")
		gitSource.Sync()
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
