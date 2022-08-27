package handler

import (
	"devxstats/model"
	"devxstats/storage"
	"fmt"
	"net/http"
	"time"
)

func AddCommits(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request for adding commit")
	events := []interface{}{
		model.Commit{
			Timestamp: time.Now(),
			Team:      "devx",
			Group:     "ABCD",
			Repo:      "devxstats",
			System:    "Github",
			User:      "tomas-mota",
		},
		model.Commit{
			Timestamp: time.Now(),
			Team:      "devx",
			Group:     "EFGH",
			Repo:      "devxstats2",
			System:    "bitbucket",
			User:      "jane doe",
		},
	}

	err := storage.DBStore.AddCommits(events)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, events)
}

func GetCommits(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request for events")
	group := "ABCD" // for now assume request is for events in this project

	events, err := storage.DBStore.GetCommits(group)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, events)
}
