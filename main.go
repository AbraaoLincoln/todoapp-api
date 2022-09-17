package main

import (
	"github.com/abraaolincoln/todoapp-api/database"
	log "github.com/abraaolincoln/todoapp-api/logger"
)

func main() {
	log.Info("Starting todoapp api...")
	database.Connect()
	database.Migrate()
}
