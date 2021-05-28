package main

import "fmt"

func main() {
	/*
		arrays são estruturas homogêneas e estáticas (só possuem o mesmo tipo de dado dentro deles
		e não variam de tamanho)
	*/
	var notas [3]float64 // nomeDoArray [tamanho]tipo
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1
	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}
	media := total / float64(len(notas))
	fmt.Printf("Média %.2f\n", media)
}
