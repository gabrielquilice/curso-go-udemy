package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma a+b=", a+b)
	fmt.Println("Subtração a-b=", a-b)
	fmt.Println("Divisão a/b=", a/b)
	fmt.Println("Divisão a*b=", a*b)
	fmt.Println("Módulo (resto da divisão)=", a%b)

	//bitwise - operações bit a bit
	fmt.Println("E =", a&b) //valores binários de a e b (11 & 10 = 10)
	fmt.Println("OU =", a|b)
	fmt.Println("XOR =", a^b)

	c := 3.0
	d := 2.0

	//outras operações
	fmt.Println("Maior =", math.Max(float64(a), float64(b)))
	fmt.Println("Menor =", math.Min(c, d))
	fmt.Println("Exponenciação =", math.Pow(c, d))
}
