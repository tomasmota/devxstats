package handler

import (
	"devxstats/storage"
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

func GetCommits(w http.ResponseWriter, r *http.Request) {
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
