package main

import "fmt"

//init é uma convenção em go; essa função será executada antes mesmo do main; serve pra inicializações 
func init() { 
	fmt.Println("Inicializando...")
}

func main() {
	fmt.Println("Função main!")
}
