package main

import "fmt"

/*
	Quando um canal é definido com buffer, não haverá operações de bloqueio;
	será esperado até que o canal inteiro seja preenchido, para aí sim bloqueado
*/
func rotina(ch chan int) {
	fmt.Println("Executou!")
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	ch <- 6
}

func main(){
	ch := make(chan int, 3)
	go rotina(ch)

	fmt.Println(<-ch)
}
