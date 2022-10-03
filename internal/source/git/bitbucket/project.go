package bitbucket

import "devxstats/internal/model"

type Page struct {
	Projects      []Project `json:"values"`
	Size          int       `json:"size"`
	IsLastPage    bool      `json:"isLastPage"`
	NextPageStart int       `json:"nextPageStart"`
	Start         int       `json:"start"`
	Limit         int       `json:"limit"`
}

type Project struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
	Key         string `json:"key"`
	ID          int    `json:"id"`
	Type        string `json:"type"`
	Public      bool   `json:"public"`
}

func (p *Project) toGroup() *model.Group {
	return &model.Group{
		Name:        p.Name,
		Key:         p.Key,
		Description: p.Description,
	}
}
