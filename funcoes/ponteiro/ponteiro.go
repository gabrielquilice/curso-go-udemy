package main

import "fmt"

//o * refere-se ao valor da variável original do ponteiro; & refere-se ao endereço de memória
func inc1(n int) {
	n++ //a incrementação aqui é feita apenas na cópia do parâmetro, não no valor original, que continua o mesmo
}

func inc2(n *int) {//aqui o parâmetro é o endereço de memória de uma variável do tipo int
	*n++ //aqui, o incremento é feito diretamente na variável original, porque o valor passado não é uma cópia, mas sim a referência de memória da variável de origem
}

func main() {
	n := 1
	inc1(n)
	fmt.Println(n)

	inc2(&n) //passando o endereço de memória da variável n
	fmt.Println(n)
}
