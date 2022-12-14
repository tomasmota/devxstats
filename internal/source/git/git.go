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
	GetGroups(ctx context.Context) ([]*model.Group, error)
	GetRepos(ctx context.Context) ([]*model.Repo, error)
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
		gc, err := github.NewClient(c.Github.Token)
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
		system, err := s.db.GetSystemByName(ctx, source.Name())
		if err != nil {
			return fmt.Errorf("error fetching system entity for %s: %w", source.Name(), err)
		}

		// GROUPS
		groups, err := source.GetGroups(ctx)
		if err != nil {
			return err
		}

		for _, g := range groups {
			g.SystemID = system.ID
			err := s.db.AddGroup(ctx, *g)
			if err != nil {
				return fmt.Errorf("error persisting group %v: %w", g.Name, err)
			}
		}
		fmt.Printf("%v: finished syncing groups\n", source.Name())

		// // REPOS
		// repos, err := source.GetRepos(ctx)
		// if err != nil {
		// 	return err
		// }

		// for _, r := range repos {
		// 	s.db.AddRepo(ctx, *r)
		// 	if err != nil {
		// 		return fmt.Errorf("error persisting repo %v: %w", r.Name, err)
		// 	}
		// }
		// fmt.Printf("%v: finished syncing repos\n", source.Name())
	}
	return nil
}
