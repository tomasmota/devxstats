package handler

import (
	"devxstats/storage"
	"fmt"
	"net/http"

	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

type CommitFilter struct {
	Team   string `schema:"team"`
	System string `schema:"system"`
	Group  string `schema:"group"`
	Repo   string `schema:"repo"`
	User   string `schema:"user"`
}

// func AddCommits(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("request for adding commit")
// 	events := []interface{}{
// 		model.Commit{
// 			Timestamp: time.Now(),
// 			Team:      "devx",
// 			Group:     "ABCD",
// 			Repo:      "devxstats",
// 			System:    "Github",
// 			User:      "tomas-mota",
// 		},
// 		model.Commit{
// 			Timestamp: time.Now(),
// 			Team:      "devx",
// 			Group:     "EFGH",
// 			Repo:      "devxstats2",
// 			System:    "bitbucket",
// 			User:      "jane doe",
// 		},
// 	}

// 	err := storage.DBStore.AddCommits(events)
// 	if err != nil {
// 		respondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, events)
// }

func GetCommits(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request for events")
	var filter CommitFilter
	err := decoder.Decode(&filter, r.URL.Query())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	events, err := storage.DBStore.GetCommits(filter.Group)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, events)
}
