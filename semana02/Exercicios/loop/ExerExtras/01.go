package main

import "fmt"

func main() {
	listaDeMercado := []string{"abacate", "sabonete", "azeite de oliva", "tomate", "banana", "macarrão", "cebola"}

	for indice := 0; indice < len(listaDeMercado); indice++ {
		fmt.Printf("%d) %s\n", indice+1, listaDeMercado[indice])
	}

}
