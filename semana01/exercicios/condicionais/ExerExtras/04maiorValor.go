package main

import(
	"fmt"
	"math"
)

func main() {
	var num1 float64
	var num2 float64
	var num3 float64

	fmt.Println("Digite um número: ")
	fmt.Scan(&num1)
	fmt.Println("Digite outro número: ")
	fmt.Scan(&num2)
	fmt.Println("Mais uma vez digite outro número: ")
	fmt.Scan(&num3)

	numTwo := math.Max(num1, num2)
	maiorValor := math.Max(numTwo, num3)

	fmt.Println(maiorValor)

	multiplo := 13

	switch {
	case multiplo == 0 :
		fmt.Printf("é um número zerp")
	case multiplo % 2 == 0 :
		fmt.Printf("É divisivel por 2")
	case multiplo % 3 == 0 :
		fmt.Printf("É divisivel por 3")
	default:
		fmt.Printf("Não é zero, não é divisivel por 2 ou por 3")
	}

}