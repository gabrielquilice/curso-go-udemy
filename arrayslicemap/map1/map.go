package main

import "fmt"

func main() {
	/*
		nomeDoMap map[tipoDaChave]tipoDoValor
		A chave do map é um valor que o identifica, como um index, e o valor do map é o que o acompanha
	*/

	//todo map tem que ser inicializado
	aprovados := make(map[int]string) //inicializando o map com o método make
	aprovados[12345678978] = "Maria"
	aprovados[12345678979] = "Pedro"
	aprovados[12345678980] = "Ana"

	fmt.Println(aprovados)

	/*
		o método range vai fazer o for percorrer as duas variáveis existentes em aprovados, levando em conta
		as variáveis cpf e nome (se quiser ignorar uma das variáveis, é só colocar _)
	*/
	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12345678980])

	delete(aprovados, 12345678980) //delete(nomeDoMap, chaveParaExcluir)
	fmt.Println("\n##########################")
	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}
}
