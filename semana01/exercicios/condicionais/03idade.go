package main

import (
	"fmt"
)

func main() {
	var idade int
	fmt.Println("Qual a sua idade? ")
	fmt.Scan(&idade)

	if idade < 18 {
		fmt.Println("Menor de idade")
	}else if (idade >= 18 && idade <= 60) {
		fmt.Println("Maior de idade")
	}else {
		fmt.Println("Idoso")
	}
}