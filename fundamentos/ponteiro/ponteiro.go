package main

import "fmt"

func main() {
	i := 1

	//Go não permite fazer operações aritméticas com ponteiros
	var p *int = nil
	/*a definição do ponteiro é feita pelo *, seguido do tipo de ponteiro que ele será
	o nil significa nulo; o ponteiro aqui no momento não aponta para lugar nenhum */

	p = &i //o & informa que irá pegar o endereço de memória da variável i

	/*
		a vaiável p só tem acesso ao endereço de memória ao qual ele aponta;
		colanco um * antes da variável (*p, por exemplo), se terá acesso ao conteúdo do endereço
	*/

	*p++
	i++

	fmt.Println("O valor é:", *p, i)
	fmt.Println("O endereço de memória é: ", p, "|",&i)

}
