package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs) //ao chegar uma requisição, na raiz da aplicação, deve passar o fileserver fs

	log.Println("Executando...")
	//o servidor irá executar e ficar esperando -- <local>:<porta>, handle
	log.Fatal(http.ListenAndServe(":3000", nil))
}
