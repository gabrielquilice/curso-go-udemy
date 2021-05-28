package main

import (
	"fmt"
)

func main() {
	var (
		C float64
		F float64
	)

	fmt.Print("Digite a temperatura em F:")
	fmt.Scanf("%f", &F) //%f aqui lê numeros
	C = (F - 32) * 5 / 9
	fmt.Println("A temperatura em C é", C)

}
