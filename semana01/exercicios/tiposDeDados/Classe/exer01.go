// Exercício #01

//     Utilizando a palavra reservada var declare uma variável numérica do tipo int.

//     Atribua um valor a essa variável.

//     Printe na tela o seu valor.

//     Repita para 3 variáveis diferentes.

package main

import (
	"fmt"
)

func main() {
	var myVar1, myVar2, myVar3 int

	myVar1 = 10
	myVar2 = 13
	myVar3 = 15

	fmt.Printf("'%v'\n", myVar1)
	fmt.Printf("'%v'\n", myVar2)
	fmt.Printf("'%v'\n", myVar3)

	var precoBanana, precoCerveja, precoAbacate, precoSaldadinho float64

	precoBanana = 1.25
	precoCerveja = 2.98
	precoAbacate = 4.59
	precoSaldadinho = 7.29

	var pesoBanana, quantidadeCerveja, pesoAbacate, quantidadeSalgadinho float64

	pesoBanana = 2.17
	quantidadeCerveja = 6
	pesoAbacate = 5.65
	quantidadeSalgadinho = 3

	valorTotal := (precoBanana * pesoBanana) + (precoCerveja * quantidadeCerveja) + (precoAbacate * pesoAbacate) + (precoSaldadinho * quantidadeSalgadinho)

	fmt.Printf("O valor total da compra deu R$ %v reais.", valorTotal)

	var name = "Ray"
	cor := `Roxo`

	fmt.Printf("Olá, meu nome é %s e gosto da cor %s", name, cor)

	a := 10 == 10
	b := 8 <= 10
	c := 25 != 10 && 20 > 50
	d :=  80 != 81
	e := 8 >= 10

	fmt.Printf("O valor de a é %v e seu tipo é %T\n", a, a)
	fmt.Printf("O valor de b é %v e seu tipo é %T\n", b, b)
	fmt.Printf("O valor de c é %v e seu tipo é %T\n", c, c)
	fmt.Printf("O valor de d é %v e seu tipo é %T\n", d, d)
	fmt.Printf("O valor de e é %v e seu tipo é %T\n", e, e)
	

}