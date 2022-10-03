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
}

func (c *client) GetRepos(ctx context.Context) ([]*model.Repo, error) {
	return []*model.Repo{{}}, nil
}
