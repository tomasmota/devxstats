package bitbucket

import (
	"context"
	"devxstats/internal/model"
	"encoding/json"
	"fmt"
)

type ProjectPage struct {
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

func (c *client) GetGroups(ctx context.Context) ([]*model.Group, error) {
	fmt.Println("fetching groups")
	var groups []*model.Group
	p := &ProjectPage{NextPageStart: 0}

	for {
		r, err := c.httpClient.Get(fmt.Sprintf("%s/projects?start=%d", c.baseURL, p.NextPageStart))
		if err != nil {
			return nil, fmt.Errorf("error fetching projects: %w", err)
		}

		defer r.Body.Close()

		err = json.NewDecoder(r.Body).Decode(p)
		if err != nil {
			return nil, fmt.Errorf("error decoding groups from api response: %w", err)
		}

		for _, p := range p.Projects {
			g := p.mapGroup()
			groups = append(groups, g)
			fmt.Println(g.Name) // remove
		}

		if p.IsLastPage {
			break
		}
	}

	return groups, nil
}

func (p *Project) mapGroup() *model.Group {
	return &model.Group{
		Name:        p.Name,
		Key:         p.Key,
		Description: p.Description,
	}
}
