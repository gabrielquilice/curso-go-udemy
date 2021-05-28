package main

import "fmt"

type curso struct {
	nome string
}

func main() {
	var coisa interface{} //funciona como um "tipo genérico"

	fmt.Println(coisa)

	coisa = 3
	fmt.Println(coisa)

	type dinamico interface{}

	var coisa2 dinamico = "Opa"
	fmt.Println(coisa2)

	coisa2 = true
	fmt.Println(coisa2)

	coisa2 = curso{"Curso Golang"}
	fmt.Println(coisa2)
}
