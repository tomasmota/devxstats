package config

import (
	"context"
	"fmt"

	"github.com/sethvargo/go-envconfig"
)

type AppConfig struct {
	Port int       `env:"PORT,default=8080"`
	Db   *DbConfig `env:",prefix=DB_"`
	Git  *GitConfig
	Cd   *CdConfig
}

type DbConfig struct {
	Host    string `env:"HOST,required"`
	Port    int    `env:"PORT,default=27017"`
	Enabled bool
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
		panic(fmt.Errorf("error parsing environment variables into config: %v", err))
	}
	c.Cd.Octopus.Enabled = true
	c.Git.Bitbucket.Enabled = false
	c.Git.Github.Enabled = true
	return c
}
