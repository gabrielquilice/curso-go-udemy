package main

import (
	"curso-go/mypkgs"
	"fmt"
	"time"
)

func oMaisRapido(url1, url2, url3 string) string {
	html := &mypkgs.Html{} 
	c1 := html.Titulo(url1)
	c2 := html.Titulo(url2)
	c3 := html.Titulo(url3)

	//Select é uma estrutura de controle específica para concorrência
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(700 * time.Millisecond):
		return "Todos os sites são lentos"
	default:
		return "Sem resposta"
	}
}

func main() {
	campeao := oMaisRapido(
		"https://www.mercadolivre.com.br",
		"https://www.youtube.com",
		"https://www.udemy.com",
	)

	fmt.Println(campeao)
}
