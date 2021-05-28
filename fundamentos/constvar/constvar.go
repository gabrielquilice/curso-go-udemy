package main

import (
	"fmt"
	m "math" // definindo um alias para um pacote
)

func main() {
	//const. - nome - tipo - valor
	const PI float64 = 3.1415

	var raio = 3.2 //(float64 é o valor padrão de números de ponto flutuante)

	//forma reduzida para declarar variável e atribuir valor
	area := PI * m.Pow(raio, 2)
	fmt.Println("A área da circunferência é", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	//1ª variável - 2ª variável - tipo das variáveis e valor da primeira - valor da segunda
	var e, f bool = true, false

	//1ª var. - 2ª var. - 3ª var - := - tipo da 1ª - tipo da 2ª - tipo da 3ª
	g, h, i := 2, false, "epa!"

	fmt.Println(a, b, c, d, e, f, g, h, i)
}
