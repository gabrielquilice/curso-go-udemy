package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	//CRIA ARQUIVO
	file, err1 := os.Create("teste.txt")
	if err1 != nil {
		return
	}
	defer file.Close()
	file.WriteString("aloo") //grava texto no arquivo

	//LÊ O ARQUIVO
	bs, err := ioutil.ReadFile("teste.txt")
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)

	//r := errors.New("erro")
	//fmt.Println(r)

}
