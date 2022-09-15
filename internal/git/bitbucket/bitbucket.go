package bitbucket

import (
	"context"
	"devxstats/model"
	"fmt"

	"devxstats/internal/util"

	"github.com/drone/go-scm/scm"
	"github.com/drone/go-scm/scm/driver/stash"
)

type BitbucketConfig struct {
	BaseUrl string
	Token   string
}

type bitbucketClient struct {
	Client *scm.Client
}

const system = "bitbucket"

func NewBitbucketClient(config *BitbucketConfig) (*bitbucketClient, error) {
	fmt.Println("creating bitbucket client, endpoint: ", config.BaseUrl)

	c, err := stash.New(config.BaseUrl)
	if err != nil {
		return nil, fmt.Errorf("an error occured while creating bitbucket client: %v", err)
	}

	c.Client = util.NewBearerHttpClient(config.Token)

	return &bitbucketClient{Client: c}, nil
}

func (c *bitbucketClient) Name() string {
	return system
}

func (c *bitbucketClient) GetOpenPullRequests(ctx context.Context) ([]*model.PullRequest, error) {
	fmt.Println("fetching bitbucket open pull requests")

	var prCount, repoCount int
	page := 1

	for {
		fmt.Println("repos loop in prs")
		opts := scm.ListOptions{
			Page: page,
			Size: 1000,
		}
		repos, res, err := c.Client.Repositories.List(ctx, opts)
		if err != nil {
			return nil, fmt.Errorf("error fetching repositories: %v", err)
		}
		if res.Status != 200 {
			return nil, fmt.Errorf("error fetching repositories, received status: %v", res.Status)
		}

		page = res.Page.Next
		repoCount += len(repos)

		for _, r := range repos {
			prs, res, err := c.Client.PullRequests.List(ctx, fmt.Sprintf("%v/%v", r.Namespace, r.Name), scm.PullRequestListOptions{Open: true})
			if err != nil {
				return nil, fmt.Errorf("error fetching repositories: %v", err)
			}
			if res.Status != 200 {
				return nil, fmt.Errorf("error fetching repositories, received status: %v", res.Status)
			}

			if len(prs) > 0 {
				fmt.Printf("%v open pull requests in repo %v\n", len(prs), fmt.Sprintf("%v/%v", r.Namespace, r.Name))
			}
			prCount += len(prs)
		}

		if res.Page.Next == 0 {
			break
		}
	}
	fmt.Printf("found %d open pull requests across %d repos\n", prCount, repoCount)

	prs := []*scm.PullRequest{{}} // TODO: fetch prs here
	return convertPullRequests(prs...), nil
}

func (c *bitbucketClient) GetCommits(ctx context.Context) ([]*model.Commit, error) {
	fmt.Println("fetching bitbucket commits")
	commits := []*scm.Commit{{}} // TODO: Fetch commits here
	return convertCommits(commits...), nil
}

func (c *bitbucketClient) GetRepositories(ctx context.Context) ([]*model.Repository, error) {
	fmt.Println("fetching bitbucket repositories")
	var allRepos []*scm.Repository
	page := 1

	for {
		opts := scm.ListOptions{
			Page: page,
			Size: 100,
		}
		repos, res, err := c.Client.Repositories.List(ctx, opts)
		if err != nil {
			return nil, fmt.Errorf("error fetching repositories: %v", err)
		}
		if res.Status != 200 {
			return nil, fmt.Errorf("error fetching repositories, received status: %v", res.Status)
		}

		page = res.Page.Next
		allRepos = append(allRepos, repos...)

		if res.Page.Next == 0 {
			break
		}
	}
	fmt.Printf("found %v repos\n", len(allRepos))
	return convertRepositories(allRepos...), nil
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
