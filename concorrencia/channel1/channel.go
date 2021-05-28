package main

import "fmt"

/* canal é um tipo; é a forma de comunicação entre as goroutines; */
func main() {
	//chan - criando um canal que envia valores (inteiros, neste caso)
	ch := make(chan int, 1)

	ch <- 1 //enviando para o canal o valor 1
	<-ch //recebendo valor do canal

	ch <- 2
	fmt.Println(<-ch) 
}
