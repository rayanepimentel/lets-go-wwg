package main

import (
	"fmt"
)

func main() {

// 	Declara uma variável;
//  Atribui o valor de um número a ela;
//  Informa se o número é positivo ou negativo.
    var num int

	fmt.Println("Digite um número positivo ou negativo.")
	fmt.Scan(&num)

	if num >= 0 {
		fmt.Println("É positivo")
	}else {
		fmt.Println("É negativo")
	}
}