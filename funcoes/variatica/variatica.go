package main

import "fmt"

//Funções variáticas são definidas como funções que possuem uma quantidade n de parâmetros de determinado tipo
func media(numeros ...float64) float64 {
	total := 0.0
	for _, valor := range numeros {
		total += valor
	}
	return total / float64(len(numeros))
}

func main() {
	fmt.Printf("Média: %.2f", media(1, 2, 3, 5))
}
