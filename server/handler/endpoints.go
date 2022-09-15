package handler

import (
	"devxstats/storage"
	"net/http"

	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

type CommitFilter struct {
	System string `schema:"system"`
	Group  string `schema:"group"`
	Repo   string `schema:"repo"`
	User   string `schema:"user"`
}

type RepoFilter struct {
	System string `schema:"system"`
	Group  string `schema:"group"`
	Name   string `schema:"repo"`
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

func GetRepositories(w http.ResponseWriter, r *http.Request) {
	var filter RepoFilter
	err := decoder.Decode(&filter, r.URL.Query())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	events, err := storage.DBStore.GetRepos(filter.Group)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, events)
}
