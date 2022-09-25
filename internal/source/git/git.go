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
	Name() string
}

func NewGitSyncer(c *config.GitConfig, db db.DB) *GitSyncer {
	syncer := &GitSyncer{}
	if c.Bitbucket.Enabled {
		bc, err := bitbucket.NewBitbucketClient(
			&bitbucket.BitbucketConfig{
				BaseUrl: c.Bitbucket.Url,
				Token:   c.Bitbucket.Token,
			})
		if err != nil {
			panic(err)
		}
		syncer.sources = append(syncer.sources, bc)
	}
	if c.Github.Enabled {
		githubClient, err := github.NewClient(
			&github.GithubConfig{
				BaseUrl: c.Github.Url,
				Token:   c.Github.Token,
			})
		if err != nil {
			panic(err)
		}
		syncer.sources = append(syncer.sources, githubClient)
	}

	syncer.db = db

	return syncer
}

func (s *GitSyncer) Sync(ctx context.Context) error {
	for _, source := range s.sources {
		// _, err = source.GetOpenPullRequests(ctx)
		// if err != nil {
		// 	return err
		// }

		// TODO: Persist PullRequests

		// REPOS
		repos, err := source.GetRepositories(ctx)
		if err != nil {
			return err
		}

		for _, r := range repos {
			s.db.AddRepo(ctx, *r)
			if err != nil {
				return err
			}
		}
		fmt.Println("finished syncing repos from", source.Name())
	}
	return nil
}
