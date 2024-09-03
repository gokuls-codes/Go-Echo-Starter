package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	db, err := sql.Open("sqlite3", "./database.sqlite3")
	if err != nil {
		log.Fatal(err)
	}

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{}) 
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"sqlite3", driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	cmd := os.Args[(len(os.Args) - 1)]
	if cmd == "up" {
		err = m.Up()
		if err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	} else if cmd == "down" {
		err = m.Down()
		if err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	} else {
		log.Fatal("Invalid command")
	}
}