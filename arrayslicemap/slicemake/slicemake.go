package main

import "fmt"

func main() {
	s := make([]int, 10) //como nenhum array foi definido, o comando make referencia um array interno, com 10 posiçoes para o slice
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20) //aqui, o slice possui as 10 posições, mas o array interno referenciado tem 20 posições
	fmt.Println(s, len(s), cap(s))
	/*
		len() - encontra o tamanho
		cap() - encontra a capacidade do slice (o tamanho do array interno)
	*/

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	/*
		foi incorporado ao slice mais posições, porém o array interno continua o mesmo, com a mesma capacidade;
		caso o tamanho do slice ultrapasse a capacidade do array interno, a própria linguagem irá referenciar um array maior, com os mesmo valores
	*/
	fmt.Println(s, len(s), cap(s))
}
