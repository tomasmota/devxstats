package bitbucket

import (
	"context"
	"devxstats/internal/model"
)

type RepoPage struct {
	Repos         []Repo `json:"values"`
	Size          int    `json:"size"`
	IsLastPage    bool   `json:"isLastPage"`
	NextPageStart int    `json:"nextPageStart"`
	Start         int    `json:"start"`
	Limit         int    `json:"limit"`
}

type Repo struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
	ID   int    `json:"id"`
}

func (c *client) GetRepos(ctx context.Context) ([]*model.Repo, error) {
	return []*model.Repo{{}}, nil
}
