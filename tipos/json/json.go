package main

import (
	"encoding/json"
	"fmt"
)

//JSON é um formato para transmissão de dados
type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	//struct para json
	p1 := produto{1, "Notebook", 1899.90, []string{"Promoção", "Eletrônicos"}}
	p1Json, _ := json.Marshal(p1) // Marshal retorna dois valores, sendo um deles um error (aqui será ignorado)
	fmt.Println(string(p1Json))

	//json para struct
	var p2 produto
	jsonString := `{"id":2, "nome":"Caneta","preco":1.50, "tags":["Papelaria", "Nacional"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Tags[1])
}
