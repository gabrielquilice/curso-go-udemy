package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("UPDATE usuarios SET nome=$1 WHERE id=$2")
	stmt.Exec("José", 1)
	stmt.Exec("Eduarda", 2)
}
