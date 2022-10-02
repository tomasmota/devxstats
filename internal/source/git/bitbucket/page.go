package bitbucket

type Page struct {
	Values        interface{} `json:"values"`
	Size          int         `json:"size"`
	IsLastPage    bool        `json:"isLastPage"`
	NextPageStart int         `json:"nextPageStart"`
	Start         int         `json:"start"`
	Limit         int         `json:"limit"`
}

func (p *Page) Next() error {
	return nil
}
