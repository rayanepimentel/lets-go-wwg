package main

import (
	"fmt"
)

type conjuntoInteiro []int

func (c conjuntoInteiro) some() int {
	var soma int
	for _, value := range c {
		soma += value
	}

	return soma
}

func (c conjuntoInteiro) media() float64 {
	soma := c.some()
	media := float64(soma) / float64(len(c))
	return media
}

func main() {

	conjunto1 := conjuntoInteiro{2, 4, 6, 8}

	soma := conjunto1.some()
	fmt.Println("soma:", soma)

	media := conjunto1.media()
	fmt.Println("média:", media)

	

}