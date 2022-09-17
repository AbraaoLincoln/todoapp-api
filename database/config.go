package database

import (
	"database/sql"
	"os"

	log "github.com/abraaolincoln/todoapp-api/logger"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

const databaseFileLocation = "./database/files/"
const devDatabase = "todoapp_dev.db"

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

func GetDatabase() *sql.DB {
	return db
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
