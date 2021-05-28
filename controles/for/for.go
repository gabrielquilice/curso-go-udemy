package main

import "fmt"

/*
	Dentro do Go, não existem outras estruturas de repetição fora o própio for; o que muda
	é a estrutura do for em cada caso, como quando não se sabe o tamanho específico do laço
*/

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\n\nMisturando")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Ímpar ")
		}
	}

	//laço infinito
	/*
		for {
			fmt.Println("Para sempre...")
			time.Sleep(time.Second) //define uma espera de 1 segundo
		}
	*/
}
