package main

import (
	"database/sql"

	_ "github.com/lib/pq" //o _ indica um import implícito; só iremos acessá-lo
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	//sql.Open("postgres", "host=% port=% user=% password=% dbname=% sslmode=disable")
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//exec(db, "CREATE DATABASE curso")
	exec(db, "DROP TABLE IF EXISTS usuarios")
	exec(db, `CREATE TABLE usuarios (
        id serial,
        nome varchar(80), 
        PRIMARY KEY (id))`)
}
