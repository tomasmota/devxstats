package config

import (
	"context"
	"errors"
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

func (c *AppConfig) validate() error {
	if c.Git.Bitbucket.Enabled {
		if c.Git.Bitbucket.Url == "" {
			return errors.New("if bitbucket source is enabled, BITBUCKET_URL must be set")
		}
		if c.Git.Bitbucket.Token == "" {
			return errors.New("if bitbucket source is enabled, BITBUCKET_TOKEN must be set")
		}
	}
	if c.Git.Github.Enabled {
		if c.Git.Github.Token == "" {
			return errors.New("if github source is enabled, GITHUB_TOKEN must be set")
		}
	}
	if c.Cd.Octopus.Enabled {
		if c.Cd.Octopus.Url == "" {
			return errors.New("if octopus source is enabled, OCTOPUS_URL must be set")
		}
		if c.Cd.Octopus.Token == "" {
			return errors.New("if octopus source is enabled, OCTOPUS_TOKEN must be set")
		}
	}
	return nil
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

	if err := c.validate(); err != nil {
		panic(fmt.Errorf("config is not valid: %w", err))
	}
	return c
}
