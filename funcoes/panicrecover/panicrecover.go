package main

import "fmt"

func main() {
	defer func() {
		str := recover() //recover interrompe o panic, devolvendo o valor chamado para a função
		fmt.Println(str)
	}()
	panic("Erro")
}
