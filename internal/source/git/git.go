package git

import (
	"context"
	"devxstats/internal/config"
	"devxstats/internal/db"
	"devxstats/internal/model"
	"devxstats/internal/source/git/bitbucket"
	"devxstats/internal/source/git/github"
	"fmt"
)

type GitSyncer struct {
	sources []GitClient
	db      db.DB
}

type GitClient interface {
	GetRepositories(ctx context.Context) ([]*model.Repo, error)
	GetGroups(ctx context.Context) ([]*model.Group, error)
	Name() string
}

func NewGitSyncer(c *config.GitConfig, db db.DB) *GitSyncer {
	syncer := &GitSyncer{}
	if c.Bitbucket.Enabled {
		bc, err := bitbucket.NewClient(c.Bitbucket.Url, c.Bitbucket.Token)
		if err != nil {
			panic(err)
		}
		syncer.sources = append(syncer.sources, bc)
	}
	if c.Github.Enabled {
		gc, err := github.NewClient(c.Github.Url, c.Github.Token)
		if err != nil {
			panic(err)
		}
		syncer.sources = append(syncer.sources, gc)
	}

	syncer.db = db

	return syncer
}

func (s *GitSyncer) Sync(ctx context.Context) error {
	for _, source := range s.sources {
		// GROUPS
		_, err := source.GetGroups(ctx)
		if err != nil {
			return err
		}

		// REPOS
		_, err = source.GetRepositories(ctx)
		if err != nil {
			return err
		}

		// for _, r := range repos {
		// 	s.db.AddRepo(ctx, *r)
		// 	if err != nil {
		// 		return err
		// 	}
		// }
		fmt.Println("finished syncing repos from", source.Name())
	}
	return nil
}
