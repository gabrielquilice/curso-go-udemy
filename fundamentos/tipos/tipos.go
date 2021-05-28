package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é:", reflect.TypeOf(32000)) //reflect.TypeOf retorna o tipo de uma variável

	//números sem sinal, só positivos: uint8, uint16, uint32, uint64 (o tamanho varia)
	var b byte = 255 //byte é um alias para uint8 (255 é o valor máximo)
	fmt.Println("O tipo da variável b é:", reflect.TypeOf(b))

	//números com sinal: int8, int16, int32, int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é:", reflect.TypeOf(i1))

	//representação inteira de um valor na tabela Unicode(int32)
	var i2 rune = 'a'
	fmt.Println("O rune de i2 é", i2)

	//números reais (float32, float64)
	var x float32 = 49.99 //float32(49.99) também é uma opção
	fmt.Println("O tipo de x é:", reflect.TypeOf(x))
	fmt.Println("O tipo do literal é:", reflect.TypeOf(49.99)) //valor padrão de literais é float64

	//tipo booleano
	bo := true
	fmt.Println("O tipo da variável bo é:", reflect.TypeOf(bo))

	//tipo string
	s1 := "Oi meu nome é Gabriel"
	fmt.Println(s1 + "!")
	fmt.Println("O tipo da variável s1 é:", reflect.TypeOf(s1))
	fmt.Println("O tamanho da string é:", len(s1)) //len() retorna o tamanho da string

	//string com múltiplas linhas
	s2 := `Olá
	me
	chamo
	Gabriel`
	fmt.Println("O tamanho da string s2 é", len(s2))

	//Go não trabalha com o tipo char; caracteres únicos são do tipo int32, segue o exemplo:
	char := 'a'
	fmt.Println("O tipo da variável char é:", reflect.TypeOf(char))

}
