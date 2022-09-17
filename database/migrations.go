package database

import (
	"github.com/abraaolincoln/todoapp-api/database/migrations"
	log "github.com/abraaolincoln/todoapp-api/logger"
)

func Migrate() {
	db := GetDatabase()

	for _, table := range migrations.GetTables() {
		_, err := db.Exec(table)

		if err != nil {
			log.Error("Couldn't create table")
			log.Error(err.Error())
		}
	}
}
