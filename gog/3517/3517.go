package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := openDb()
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("SELECT 1")
	if err != nil {
		log.Fatal(err)
	}
}

func openDb() *sql.DB {
	db, err := sql.Open("sqlite3", "./mockdb.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
