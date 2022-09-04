package git

import (
	"devxstats/internal/config"
	"devxstats/internal/git/bitbucket"
	"devxstats/internal/git/github"
	"devxstats/model"
	"flag"
)

type GitSyncer struct {
	sources []GitClient
}

type GitClient interface {
	GetCommits() ([]*model.Commit, error)
	GetOpenPullRequests() ([]*model.PullRequest, error)
}

func NewGitSyncer(c *config.GitConfig) *GitSyncer {
	githubEnabled := flag.Bool("github", false, "Set to true to enable github source")
	bitbucketEnabled := flag.Bool("bitbucket", false, "Set to true to enable bitbucket source")
	flag.Parse()

	syncer := &GitSyncer{}
	if *bitbucketEnabled {
		// Add sources based on configuration
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
	if *githubEnabled {
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

func (git *GitSyncer) Sync() error {
	for _, source := range git.sources {
		_, err := source.GetCommits()
		if err != nil {
			return err
		}

		// TODO: Persist Commits

		_, err = source.GetOpenPullRequests()
		if err != nil {
			return err
		}

		// TODO: Persist PullRequests
	}
	return nil
}
