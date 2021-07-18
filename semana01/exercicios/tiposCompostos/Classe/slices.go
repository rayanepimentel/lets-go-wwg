package main 

import (
	"fmt"
)

func main() {
	// Declare duas slices de int e preencha 6 posições de cada uma.
	// Some as slices, formando uma 3ª slice.
	// Printe as três.
	var num1 = []int{1, 3, 7, 9, 11, 13}

	var num2 = []int{2, 4, 6, 8, 10, 12}

	num3 := append(num1, num2...)

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
}