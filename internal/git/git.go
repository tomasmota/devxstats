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
	git := &GitSyncer{}
	// Add sources based on configuration
	git.sources = append(git.sources, bitbucket.GetClient("bitbucket.com")) //TODO: pass in config params
	git.sources = append(git.sources, github.GetClient("github.com"))       //TODO: pass in config params
	return git
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
