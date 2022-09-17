package main

import (
	"github.com/abraaolincoln/todoapp-api/controllers"
	"github.com/abraaolincoln/todoapp-api/database"
	log "github.com/abraaolincoln/todoapp-api/logger"
	"github.com/abraaolincoln/todoapp-api/router"
	"net/http"
)

func main() {
	log.Info("Starting todoapp api...")
	startDatabase()
	startHttpServer()
}

func startDatabase() {
	database.Connect()
	database.Migrate()
}

func startHttpServer() {
	mux := router.NewRestMux()
	controllers.RegisterProjectController(mux)

	http.ListenAndServe(":8080", mux)
}
