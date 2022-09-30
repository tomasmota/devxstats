package config

import (
	"context"
	"flag"
	"fmt"

	"github.com/sethvargo/go-envconfig"
)

type AppConfig struct {
	Port int `env:"PORT,default=8080"`
	Git  *GitConfig
	Cd   *CdConfig
}

type GitConfig struct {
	Github struct {
		Token   string `env:"GITHUB_TOKEN"`
		Url     string `env:"GITHUB_URL"`
		Enabled bool
	}
	Bitbucket struct {
		Url     string `env:"BITBUCKET_URL"`
		Token   string `env:"BITBUCKET_TOKEN"`
		Enabled bool
	}
}

type CdConfig struct {
	Octopus struct {
		Token   string `env:"OCTOPUS_TOKEN"`
		Url     string `env:"OCTOPUS_URL"`
		Enabled bool
	}
}

func Load(ctx context.Context) *AppConfig {
	c := &AppConfig{}
	if err := envconfig.Process(ctx, c); err != nil {
		panic(fmt.Errorf("error parsing environment variables into config: %w", err))
	}

	githubF := flag.Bool("github", false, "Set to true to enable github source")
	bitbucketF := flag.Bool("bitbucket", false, "Set to true to enable bitbucket source")
	octopusF := flag.Bool("octopus", false, "Set to true to enable octopus source")

	flag.Parse()
	c.Git.Github.Enabled = *githubF
	c.Git.Bitbucket.Enabled = *bitbucketF
	c.Cd.Octopus.Enabled = *octopusF

	return c
}
