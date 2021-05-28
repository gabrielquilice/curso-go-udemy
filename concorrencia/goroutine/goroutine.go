package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (interação %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// fale("Maria", "oi", 3) processo em série
	// fale("João", "tchau", 1)

	/*
		Criação de uma goroutine
		- colocando go na frente do comando, o compilador irá executar executá-lo de forma
		  independente (algo semelhante a uma thread) 

		Quando uma goroutine é criada, a função atribuida a ela vai ser executada em concorrência
		com a função principal (main) 
	*/
	go fale("Maria", "oi", 10)
	fale("João", "tchau", 5)
}
