package main

import (
	"context"
	"devxstats/internal/cd"
	"devxstats/internal/config"
	"devxstats/internal/git"
	"devxstats/server"
	"devxstats/storage"
	"flag"
	"fmt"
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
	// env := os.Getenv("ENVIRONMENT")
	// if env == "dev" {
	// 	godotenv.Load()
	// }

	c := &config.AppConfig{}
	if err := envconfig.Process(ctx, c); err != nil {
		panic(fmt.Errorf("error parsing environment variables into config: %v", err))
	}
	var testb bool
	flag.BoolVar(&testb, "github", false, "Set to true to enable github source")
	flag.BoolVar(&c.Enabled, "github", false, "Set to true to enable github source")
	fmt.Println("after flag")
	// flag.BoolVar(c.Git.Bitbucket.Enabled, "bitbucket", false, "Set to true to enable bitbucket source")
	// flag.BoolVar(c.Cd.Octopus.Enabled, "octopus", false, "Set to true to enable octopus source")
	flag.Parse()
	fmt.Println("after parse")

	app := &server.App{}

	storage.InitializeDB(ctx, c.Db)
	go syncSources(*c)
	app.InitializeRoutes()
	app.Run(fmt.Sprintf(":%d", c.Port))
}
