package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func conectar() (db *sql.DB) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=cursogo")
	if err != nil {
		log.Fatal(err)
	}
	return
}
