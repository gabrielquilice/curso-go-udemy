package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha\n")

	fmt.Println("\nNova")
	fmt.Println("linha")

	x := 3.1415

	//fmt.Println("O valor de x é", x)

	//variável convertida para string
	xs := fmt.Sprint(x)
	//string pode ser concatenada usando o +
	fmt.Println("O valor de x é " + xs)

	fmt.Printf("O valor de x é %.2f.", x)

	a := 1
	b := 1.9999
	c := false
	d := "epa"

	/*
		%d - inteiros
		%f - ponto flutuante (com todas as casas)
		%.1f - ponto flutuante com uma casa decimal
		%t - valor booleano
		%s - string
		%v - geralmente consegue formatar qualquer tipo
	*/
	fmt.Printf("\n\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
