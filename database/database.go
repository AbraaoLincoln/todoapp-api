package database

import (
	"database/sql"
	"os"

	"github.com/abraaolincoln/todoapp-api/database/migrations"
	log "github.com/abraaolincoln/todoapp-api/logger"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

const databaseFileLocation = "./database/files/"
const devDatabase = "todoapp_dev.db"

func GetDatabase() *sql.DB {
	return db
}

func Connect() {
	createDatabaseIfNotExists()

	log.Info("Connection to database...")
	sqlite3db, err := sql.Open("sqlite3", databaseFileLocation+devDatabase)

	if err != nil {
		log.Error("Couldn't connect to database")
		log.FatalError(err.Error())
	}

	db = sqlite3db
	log.Info("Successfully connected to database")
}

func createDatabaseIfNotExists() {
	_, err := os.Stat(databaseFileLocation + devDatabase)

	if os.IsNotExist(err) {
		log.Info("Database " + devDatabase + " doesn't exists")
		createDatabaseFile(devDatabase)
		return
	}

	log.Info("Database " + devDatabase + " already exists")
}

func createDatabaseFile(databaseName string) {
	log.Info("Creating database file")
	file, err := os.Create(databaseFileLocation + databaseName)

	if err != nil {
		log.Info(err.Error())
		log.FatalError("Couldn't create database file")
	}

	file.Close()
	log.Info("Successfully created database file")
}

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
