package main

import "fmt"

//criando um novo tipo
type produto struct {
	nome     string //identificador e tipo dos atributos do struct
	preco    float64
	desconto float64
}

//método com receiver - define antes do nome da função a estrutura que ela terá; no caso se usou um alias para o tipo produto
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 = produto{
		nome:     "lápis",
		preco:    1.70,
		desconto: 0.20,
	}

	fmt.Println(produto1, produto1.precoComDesconto())

	produto2 := produto{"Notebook", 1899.90, 5.00}
	fmt.Println(produto2, produto2.precoComDesconto())
}
