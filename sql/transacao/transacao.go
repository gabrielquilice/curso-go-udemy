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

	tx, _ := db.Begin() //db.Begin() inicia uma série de comandos que ainda não foram para o banco
	stmt, _ := tx.Prepare("INSERT INTO usuarios(id, nome) VALUES ($1, $2)")
	stmt.Exec(200, "Beatriz")
	stmt.Exec(201, "Carlos")
	_, err = stmt.Exec(1, "Tiago")
	if err != nil {
		tx.Rollback() //tx.Rollback() desfaz todas as operações
		log.Fatal(err)
	}
	tx.Commit() //tx.Commit() termina a série iniciada com o db.Begin() e lança os dados no banco
}
