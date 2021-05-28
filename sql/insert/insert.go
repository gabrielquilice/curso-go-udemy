package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	//sql.Open("postgres", "host=% port=% user=% password=% dbname=% sslmode=disable")
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("INSERT INTO usuarios(nome) values($1)")
	stmt.Exec("Maria")
	stmt.Exec("João")

	res, _ := stmt.Exec("Pedro")
	
	//id, _ := res.LastInsertId()
	//fmt.Println(id)

	linhas, _ := res.RowsAffected()
	fmt.Println("Linhas afetadas:", linhas)
}
