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
	app.Router.HandleFunc("/events", handler.AddEvent).Methods("POST")
	// app.Router.HandleFunc("/products", handler.CreateProduct).Methods("POST")
}

func (app *App) Run(address string) {
	log.Fatal(http.ListenAndServe(address, app.Router))
}
