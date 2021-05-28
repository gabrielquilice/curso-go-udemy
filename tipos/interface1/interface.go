package main

import "fmt"

//a interface define um conjunto de métodos que um tipo deve ter para implementar a interface 
type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

//tod interface é implementada de maneira implícita
func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	var coisa imprimivel = pessoa{"Roberto", "Silva"}

	fmt.Println(coisa.toString())
	imprimir(coisa)

	coisa = produto{"Calça", 79.50}
	coisa.toString()
	imprimir(coisa)

	imprimir(produto{"Camiseta", 82.50})
}
