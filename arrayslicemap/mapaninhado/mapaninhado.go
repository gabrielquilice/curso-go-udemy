package main

import "fmt"

func main() {
	//aqui o valor do map é outro map
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela":  100.33,
			"Guilherme": 100.34,
		},
		"J": {
			"José": 500.36,
		},
		"P": {
			"Pedro": 10000.84,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}

}
