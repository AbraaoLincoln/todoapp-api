package database

import (
	"database/sql"

	"github.com/abraaolincoln/todoapp-api/database/migrations"
	log "github.com/abraaolincoln/todoapp-api/logger"
)

func Migrate() {
	log.Info("Migrating database")
	db := GetDatabase()
	migrateTables(db)
}

func migrateTables(db *sql.DB) {
	for _, table := range migrations.GetTables() {
		_, err := db.Exec(table)

		if err != nil {
			log.Error("Couldn't create table")
			log.Error(err.Error())
		}
	}
}
