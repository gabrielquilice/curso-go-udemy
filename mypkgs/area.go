package mypkgs

import "math"

type Area struct {}

const Pi = 3.1415 //constante pública

func (a *Area) Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

func (a *Area) Rect(base, altura float64) float64 {
	return base * altura
}

//é uma funcao privada
func (a *Area) _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
