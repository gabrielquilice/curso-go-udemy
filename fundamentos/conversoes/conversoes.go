package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println("Nota final:", notaFinal)

	//int para string
	fmt.Println("Teste " + strconv.Itoa(123))

	/*string para int
	este método abaixo rtorna dois valor, um número ou um erro (caso o valor passado seja
	inválido); o _ no código, define estar ignorando o caso de erro
	*/
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro!")	
	}

}
