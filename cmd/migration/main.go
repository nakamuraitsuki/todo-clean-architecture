package main

import (
	"log"
	"practice/infrastructure/sqlite3"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sqlx.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}

	err = sqlite3.MigrateTodo(db)
	if err != nil {
		log.Fatal(err)
	}
}
