package handler

import (
	"devxstats/internal/storage"
	"net/http"

	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

type RepoFilter struct {
	System string `schema:"system"`
	Group  string `schema:"group"`
	Name   string `schema:"repo"`
}

func GetRepositories(w http.ResponseWriter, r *http.Request) {
	var filter RepoFilter
	err := decoder.Decode(&filter, r.URL.Query())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	events, err := storage.DBStore.GetRepos(123) // get a mapping from groupname to id, or just search by name? Should search be done across all systems by default?
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, events)
}
