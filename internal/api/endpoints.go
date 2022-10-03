package api

import (
	"devxstats/internal/model"
	"net/http"

	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

func (s *HTTPServer) GetSystems(w http.ResponseWriter, r *http.Request) {
	systems, err := s.db.GetSystems(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, systems)
}

func (s *HTTPServer) GetGroups(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func (s *HTTPServer) GetRepositories(w http.ResponseWriter, r *http.Request) {
	var filter model.RepoFilter
	err := decoder.Decode(&filter, r.URL.Query())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	//TODO: pass filter in here
	repos, err := s.db.GetRepo(r.Context(), 123) // get a mapping from groupname to id, or just search by name? Should search be done across all systems by default?
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, repos)
}
