package main

import (
	"fmt"
	"math"
)

type Calcular interface {
	CalcularArea() float64
}


type Quadrado struct {
	lado float64
}

func (q Quadrado) CalcularArea() float64 {
	area := math.Pow(q.lado, 2)
	return area
}

type Circulo struct {
	raio float64
}

func (c Circulo) CalcularArea() float64 {
	area := math.Pi * math.Pow(c.raio, 2)
	return area
}

func ReporteCalculo(calc Calcular) {
	area := calc.CalcularArea()
	fmt.Printf("O calculo da area dessa forma é %f\n", area)
}
func main() {

	//https://play.golang.org/p/X5iuGP_q_OX
	//https://play.golang.org/p/U0CuH3xrO6I

	circulo1 := Circulo {
		raio: 20,
	}

	quadrado1 := Quadrado {
		lado: 10,
	}


	ReporteCalculo(circulo1)
	ReporteCalculo(quadrado1)

	

	
	
}