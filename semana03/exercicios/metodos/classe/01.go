package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	raio float64
}

//metodo area

func (c Circulo) CalcularArea() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

//método perimetro

func (c Circulo) CalcularPerimetro() float64 {
	return 2 * math.Pi * c.raio
}
func main() {

	circuloCalc := Circulo {
		raio: 3,
	}

	areaDoCirculo := circuloCalc.CalcularArea()
	perimetroDoCirculo := circuloCalc.CalcularPerimetro()

	fmt.Printf("Círculo \n\traio: %.2f\n\tárea: %2.f\n\tperímetro: %.2f", circuloCalc.raio, areaDoCirculo, perimetroDoCirculo )
}