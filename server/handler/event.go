package handler

import (
	"devxstats/model"
	"devxstats/storage"
	"fmt"
	"net/http"
	"time"
)

func AddEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request for adding event")
	event := model.CommitEvent{
		Timestamp: time.Now(),
		Project:   "ABCD",
		Repo:      "devxstats",
	}

	err := storage.DBStore.AddEvent(event)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, event)
}

func GetEvents(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request for events")
	projectName = "ABCD" // for now assume request is for events in this project

	err := storage.DBStore.GetEvents(event)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, event)
}
