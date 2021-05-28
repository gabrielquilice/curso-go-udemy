package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} //array
	s1 := []int{1, 2, 3}  //slice - estrutura de tamanho variável; é um trecho de um array

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), " - ", reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}
	s2 := a2[1:3] // irá passar os valores entre o índice 1 e 3 (excluindo o valor do index 3) do array a2
	fmt.Println(a2, s2)

	s3 := a2[:2] //vai do início do array até o índice 2, exclusive
	fmt.Println(a2, s3)

	s4 := s2[:1]
	fmt.Println(s2, s4)

}
