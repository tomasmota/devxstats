package bitbucket

import (
	"context"
	"devxstats/internal/model"
	"devxstats/internal/util"
	"fmt"
	"net/http"

	"github.com/drone/go-scm/scm"
)

type BitbucketProject struct {
	Description string `json:"description"`
	Namespace   string `json:"namespace"`
	Avatar      string `json:"avatar"`
	Scope       string `json:"scope"`
	Name        string `json:"name"`
	Key         string `json:"key"`
	ID          int    `json:"id"`
	Type        string `json:"type"`
	Public      bool   `json:"public"`
}

const (
	apiPath = "/rest/api/1.0"
)

type client struct {
	baseURL    string
	token      string
	httpClient *http.Client
}

const system = "bitbucket"

func NewClient(baseURL string, token string) (*client, error) {
	fmt.Println("creating %v client, endpoint: ", system, baseURL)

	c := &client{
		baseURL:    fmt.Sprintf("%s%s", baseURL, apiPath),
		token:      token,
		httpClient: util.NewBearerHttpClient(token),
	}
	_, err := c.GetGroups(context.Background()) // use GetGroups as a test call
	if err != nil {
		return nil, fmt.Errorf("error creating %v client: %w", system, err)
	}

	return c, nil
}

func (c *client) Name() string {
	return system
}

func (c *client) GetGroups(ctx context.Context) ([]*model.Group, error) {
	fmt.Println("fetching groups")
	groups := []*model.Group{}
	// pagedGroups := &Page{Values: groups}

	// r, err := c.httpClient.Get(fmt.Sprintf("%s/projects", c.baseURL))
	// if err != nil {
	// 	fmt.Errorf("error fetching projects: %w", err)
	// }

	// defer r.Body.Close()
	// err = json.NewDecoder(r.Body).Decode(pagedGroups)
	// if err != nil {
	// 	return nil, fmt.Errorf("error decoding groups from api response: %w", err)
	// }

	return groups, nil
}

func (c *client) GetRepositories(ctx context.Context) ([]*model.Repo, error) {
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
			return nil, fmt.Errorf("error fetching repositories: %w", err)
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

func (c *client) GetOpenPullRequests(ctx context.Context) ([]*model.PullRequest, error) {
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
			return nil, fmt.Errorf("error fetching repositories: %w", err)
		}
		if res.Status != 200 {
			return nil, fmt.Errorf("error fetching repositories, received status: %v", res.Status)
		}

		page = res.Page.Next
		repoCount += len(repos)

		for _, r := range repos {
			prs, res, err := c.Client.PullRequests.List(ctx, fmt.Sprintf("%v/%v", r.Namespace, r.Name), scm.PullRequestListOptions{Open: true})
			if err != nil {
				return nil, fmt.Errorf("error fetching pull requests: %w", err)
			}
			if res.Status != 200 {
				return nil, fmt.Errorf("error fetching pull requests, received status: %v", res.Status)
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

func convertRepositories(from ...*scm.Repository) []*model.Repo {
	var to []*model.Repo
	for _, r := range from {
		to = append(to, &model.Repo{
			Name: r.Name,
		})
	}
	return to
}

func convertPullRequests(from ...*scm.PullRequest) []*model.PullRequest {
	// TODO: Implement
	return []*model.PullRequest{{}}
}
