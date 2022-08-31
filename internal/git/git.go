package git

import (
	"devxstats/internal/git/bitbucket"
	"devxstats/internal/git/github"
	"devxstats/model"
)

type GitSyncer struct {
	sources []GitClient
}

type GitClient interface {
	GetCommits() ([]*model.Commit, error)
	GetOpenPullRequests() ([]*model.PullRequest, error)
}

func NewGitSyncer( /*configuration of sources will somehow get injected into this method*/ ) *GitSyncer {
	syncer := &GitSyncer{}
	// Add sources based on configuration
	bc, err := bitbucket.NewBitbucketClient("https://dcgit.dac.local")
	if err != nil {
		panic(err)
	}
	syncer.sources = append(syncer.sources, bc)

	gc, err := github.NewGithubClient("https://github.com")
	if err != nil {
		panic(err)
	}
	syncer.sources = append(syncer.sources, gc)

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
