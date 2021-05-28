package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} //definindo um array sem o tamanho especificado; nesse caso, quem conta o tamanho é o compilador
	for i, numero := range numeros {
		fmt.Printf("%d) %d\n\n", i+1, numero)
	}

	for _, num := range numeros {
		fmt.Println(num)
	}
}
