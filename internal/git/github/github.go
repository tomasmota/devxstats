package github

import (
	"context"
	"devxstats/model"
	"fmt"
	"net/http"

	"github.com/drone/go-scm/scm"
	"github.com/drone/go-scm/scm/driver/github"
	"github.com/drone/go-scm/scm/transport"
)

type GithubConfig struct {
	BaseUrl string
	Token   string
}

type githubClient struct {
	Client *scm.Client
}

const system = "github"

// Creates a new client from *GithubConfig
func NewClient(config *GithubConfig) (*githubClient, error) {
	fmt.Println("creating github client")
	var c *scm.Client
	var err error
	if config.BaseUrl != "" {
		c, err = github.New(config.BaseUrl)
		if err != nil {
			return nil, err
		}
	} else {
		c = github.NewDefault()
	}

	c.Client = &http.Client{
		Transport: &transport.BearerToken{
			Token: config.Token,
		},
	}

	return &githubClient{Client: c}, nil
}

func (c *githubClient) Name() string {
	return system
}

func (c *githubClient) GetRepositories(ctx context.Context) ([]*model.Repository, error) {
	fmt.Println("Fetching github open pull requests")
	repos := []*scm.Repository{{}} // TODO: fetch prs here
	return convertRepositories(repos...), nil
}

func (c *githubClient) GetOpenPullRequests(ctx context.Context) ([]*model.PullRequest, error) {
	fmt.Println("Fetching github open pull requests")
	prs := []*scm.PullRequest{{}} // TODO: fetch prs here
	return convertPullRequests(prs...), nil
}

func (c *githubClient) GetCommits(ctx context.Context) ([]*model.Commit, error) {
	fmt.Println("Fetching github commits")

	repos, res, err := c.Client.Repositories.List(ctx, scm.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("error fetching repositories: %v", err)
	}
	if res.Status != 200 {
		return nil, fmt.Errorf("error fetching repositories, received status: %v", res.Status)
	}
	for _, r := range repos {
		commits, res, err := c.Client.Git.ListCommits(ctx, fmt.Sprintf("%v/%v", r.Namespace, r.Name), scm.CommitListOptions{Size: 1000})
		if err != nil {
			return nil, fmt.Errorf("error fetching repositories: %v", err)
		}
		if res.Status != 200 {
			return nil, fmt.Errorf("error fetching repositories, received status: %v", res.Status)
		}
		fmt.Printf("repo %v contains %v commits\n", fmt.Sprintf("%v/%v", r.Namespace, r.Name), len(commits))
	}

	commits := []*scm.Commit{{}} // TODO: Fetch commits here
	return convertCommits(commits...), nil
}

func convertPullRequests(from ...*scm.PullRequest) []*model.PullRequest {
	// TODO: Implement
	return []*model.PullRequest{{}}
}

func convertCommits(from ...*scm.Commit) []*model.Commit {
	// TODO: Implement
	return []*model.Commit{{}}
}

func convertRepositories(from ...*scm.Repository) []*model.Repository {
	// TODO: Implement
	var to []*model.Repository
	for _, r := range from {
		to = append(to, &model.Repository{
			System: system,
			Group:  r.Namespace,
			Name:   r.Name,
		})
	}
	return to
}
