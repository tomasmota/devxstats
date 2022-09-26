package api

import (
	"devxstats/internal/db"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	Router *mux.Router
	db     db.DB
}

func NewHTTPServer(db db.DB) *HTTPServer {
	s := &HTTPServer{db: db}
	s.Router = mux.NewRouter()
	s.Router.HandleFunc("/repos", s.GetRepositories).Methods("GET")
	s.Router.HandleFunc("/systems", s.GetSystems).Methods("GET")
	return s
}

func (s *HTTPServer) Run(address string) {
	fmt.Println("initializing http server")
	log.Fatal(http.ListenAndServe(address, s.Router))
}
