package main

import (
	"fmt"
)

func main() {
	//var quilometros int8
	var quilometros int16
	quilometros = 150

	fmt.Println(quilometros)
}

// 1) Descubra por que não compila
//R: int8 signed  8-bit integers (-128 to 127) e 150 é maior que 127
// 2) Erros de compilação nos ajudam a compreender o que precisamos consertar em nosso 
// código. O que o erro ./prog.go:9:14: constant 150 overflows int8 nos indica?
//R: indica que 150 "transbordou" int8
// 3) Conserte o erro de compilação e faça a quantidade de quilômetros ser imprimida 
// na tela
// R: podemos usar int mesmo, mas seria algo como 'geralzação'. 
// Podemos usar int16