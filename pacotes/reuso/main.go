package main

import (
	"curso-go/mypkgs"
	"fmt"
)

func main() {
	area := &mypkgs.Area{}
	fmt.Println(area.Circ(6.0))
	fmt.Println(area.Rect(5.0, 2.0))
	//fmt.Println(area._TrianguloEq(5.0, 2.0))
}
