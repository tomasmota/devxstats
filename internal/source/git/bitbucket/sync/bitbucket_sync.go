package bitbucket_sync

import (
	"devxstats/internal/db"
	"devxstats/internal/util"
	"fmt"
	"net/http"
)

const apiPath = "/rest/api/1.0"

type BitbucketSyncer struct {
	httpClient *http.Client
	baseURL    string
	db         db.DB
}

func NewBitbucketSyncer(baseURL string, token string, db db.DB) (*BitbucketSyncer, error) {
	fmt.Println("creating bitbucket syncer")
	// url, err := url.Parse(baseURL)
	// if err != nil {
	// 	return nil, fmt.Errorf("an error occured while parsing bitbucket url: %w", err)
	// }

	syncer := &BitbucketSyncer{
		baseURL:    fmt.Sprintf("%s%s", baseURL, apiPath),
		httpClient: util.NewBearerHttpClient(token),
		db:         db,
	}
	return syncer, nil
}
