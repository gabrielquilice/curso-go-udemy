package main

import "fmt"

func main() {
	//inicializando o map com atribuição de valores
	funcsESalarios := map[string]float64{
		"José":  1000.00,
		"Ana":   16900.45,
		"Pedro": 4000.00,
	}
	//adicionando um novo registro
	funcsESalarios["Rafael"] = 1350.00
	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
