package main

import (
	"curso-go/mypkgs"
	"fmt"
)

//Multiplexar nada mais é que juntar as infomações de canais diferentes em um único canal

//Encaminha os dados do canal de origem para o canal de destino
func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	html := &mypkgs.Html{}
	c := juntar(
		html.Titulo("https://www.google.com", "https://www.youtube.com"),
		html.Titulo("https://www.udemy.com", "https://www.mercadolivre.com.br"),
	)

	fmt.Println("Os dois mais rápidos: ")
	fmt.Println(<-c, " ###### ", <-c)
	fmt.Println("\nOs dois mais lentos: ")
	fmt.Println(<-c, " ###### ", <-c)
}
