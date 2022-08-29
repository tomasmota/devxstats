package sources

import (
	"devxstats/model"
	"fmt"
)

type gitSources struct {
	sources []GitSource
}

type GitSource interface {
	GetCommits() ([]model.Commit, error)
	GetOpenPullRequests() ([]model.PullRequest, error)
}

func NewGit( /*configuration of sources will somehow get injected into this method*/ ) *gitSources {
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
		commits, err := source.GetCommits()
		if err != nil {
			return err
		}
		fmt.Printf("Commits: %v\n", len(commits)) //TODO: Persist to database

		openPullRequests, err := source.GetOpenPullRequests()
		if err != nil {
			return err
		}
		fmt.Printf("Open Pull Requests: %v\n", len(openPullRequests)) //TODO: Persist to database
	}
	return nil
}
