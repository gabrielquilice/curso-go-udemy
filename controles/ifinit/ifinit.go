package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano()) //pega o nano segundo da execução do código
	r := rand.New(s)                           //sempre que executar o código será gerado um número novo a partir da fonte s

	return r.Intn(10) //gera um número aleatório até dez
}

func main() {
	//iniciando a variável i dentro do if
	if i := numeroAleatorio(); i > 5 {
		fmt.Println("Ganhou!")
	} else {
		fmt.Println("Perdeu!")
	}
}
