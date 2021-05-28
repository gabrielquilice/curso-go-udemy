package main

import (
	"fmt"
	"time"
)

func main() {
	//os operadore relacionais retornam como resultado true ou false
	fmt.Println("Strings:", "A" == "A")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)
	fmt.Println("<=", 3 <= 2)
	fmt.Println(">=", 3 >= 2)

	d1 := time.Unix(0, 0) //definição de um dado tipo data
	d2 := time.Unix(0, 0)
	fmt.Println("Datas iguais: ", d1.Equal(d2))

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"João"}
	p2 := Pessoa{"João"}
	fmt.Println("Pessoas:", p1 == p2)
}
