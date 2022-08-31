package git

import (
	"devxstats/model"
	"devxstats/pkg/bitbucket"
	"devxstats/pkg/github"
)

type gitSources struct {
	sources []GitClient
}

type GitClient interface {
	GetCommits() ([]model.Commit, error)
	GetOpenPullRequests() ([]model.PullRequest, error)
}

func NewGitSources( /*configuration of sources will somehow get injected into this method*/ ) *gitSources {
	git := &gitSources{}
	// if config.contains("bitbucket") {
	git.sources = append(git.sources, bitbucket.GetClient("bitbucket.com")) //TODO: pass in config params
	// }
	// if config.contains("github") {
	git.sources = append(git.sources, github.GetClient("github.com")) //TODO: pass in config params
	// }
	return git
}

func (git *gitSources) Sync() error {
	for _, source := range git.sources {
		_, err := source.GetCommits()
		if err != nil {
			return err
		}

		_, err = source.GetOpenPullRequests()
		if err != nil {
			return err
		}
	}
	return nil
}
