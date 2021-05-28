package main

import "fmt"

//você pode definir nomes para os retornos de uma função
func trocar(p1, p2 int) (segundo int, primeiro int) {
	segundo = p2
	primeiro = p1
	return //quando você atribui o valor aos retornos, não é necessário colocar nada em return (retorno limpo)
}

func main() {
	x, y := trocar(2, 3)
	fmt.Println(x, " - ", y)
}
