package git

import (
	"context"
	"devxstats/internal/config"
	"devxstats/internal/git/bitbucket"
	"devxstats/internal/git/github"
	"devxstats/model"
	"devxstats/storage"
)

type GitSyncer struct {
	sources []GitClient
}

type GitClient interface {
	GetCommits(ctx context.Context) ([]*model.Commit, error)
	GetOpenPullRequests(ctx context.Context) ([]*model.PullRequest, error)
	GetRepositories(ctx context.Context) ([]*model.Repository, error)
}

func NewGitSyncer(c *config.GitConfig) *GitSyncer {
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

	return syncer
}

func (git *GitSyncer) Sync(ctx context.Context) error {
	for _, source := range git.sources {
		_, err := source.GetCommits(ctx)
		if err != nil {
			return err
		}

		// TODO: Persist Commits

		// _, err = source.GetOpenPullRequests(ctx)
		// if err != nil {
		// 	return err
		// }

		// TODO: Persist PullRequests

		repos, err := source.GetRepositories(ctx)
		if err != nil {
			return err
		}

		storage.DBStore.AddRepos(ctx, repos)
		// TODO: Persist Repositories
	}
	return nil
}
