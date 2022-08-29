package sources

import (
	"devxstats/model"
)

type gitSources struct {
	sources []GitSource
}

type GitSource interface {
	GetCommits() ([]model.Commit, error)
	GetOpenPullRequests() ([]model.PullRequest, error)
}

func NewGitSources( /*configuration of sources will somehow get injected into this method*/ ) *gitSources {
	git := &gitSources{}
	// if config.contains("github") {
	git.sources = append(git.sources, newGithubSource())
	// }
	// if config.contains("bitbucket") {
	git.sources = append(git.sources, newBitbucketSource())
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
