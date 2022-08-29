package sources

import "devxstats/model"

type git struct{}

type GitSource interface {
	GetCommits() ([]model.Commit, error)
	GetOpenPullRequests() ([]model.PullRequest, error)
}

func getSources( /*configuration of sources will somehow get injected into this method*/ ) GitSource {

	// var sources []*GitSource

	return BitbucketSource{baseUrl: "asdf"}

	// if config.contains("github") {
	////// sources = append(sources, newGithubSource())
	// }
	// if config.contains("bitbucket") {
	////// sources = append(sources, newBitbucketSource())
	// }
	// return sources
}

func (git *git) Fetch() {
	return
}
