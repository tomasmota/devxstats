package main

import (
	"devxstats/server"
	"devxstats/storage"
)

func main() {
	app := &server.App{}

	storage.InitializeDB()
	app.InitializeRoutes()
	app.Run(":8080")
}
