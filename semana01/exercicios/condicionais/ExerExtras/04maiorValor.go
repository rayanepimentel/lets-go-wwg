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


}