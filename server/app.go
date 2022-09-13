package server

import (
	"devxstats/server/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (app *App) InitializeRoutes() {
	app.Router = mux.NewRouter()
	app.Router.HandleFunc("/commits", handler.GetCommits).Methods("GET")
	app.Router.HandleFunc("/repos", handler.GetRepositories).Methods("GET")
}

func (app *App) Run(address string) {
	log.Fatal(http.ListenAndServe(address, app.Router))
}
