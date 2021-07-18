package main

import (
	"fmt"
)

func main() {
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