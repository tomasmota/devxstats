package handler

import (
	"devxstats/model"
	"devxstats/storage"
	"fmt"
	"net/http"
	"time"
)

// TODO: Make plural
func AddCommit(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request for adding commit")
	event := model.Commit{
		Timestamp: time.Now(),
		Group:     "ABCD",
		Repo:      "devxstats",
		System:    "Github",
		User:      "tomas-mota",
	}

	err := storage.DBStore.AddCommit(event)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, event)
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
