package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch { //nesse caso o switch irá procurar o primeiro case verdadeiro; poderia colocar "switch true"
	case t.Hour() < 12:
		fmt.Println("Bom dia!")
	case t.Hour() < 18:
		fmt.Println("Boa tarde!")
	default:
		fmt.Println("Boa noite")
	}
}
