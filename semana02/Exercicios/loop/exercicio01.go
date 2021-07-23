package main

import "fmt"

func main() {
  
  //No Exercício #06 da seção "Exercícios", usamos for range para percorrer uma slice de string que representava 
  //uma lista de itens a comprar no mercado. Agora, resolva o mesmo exercício usando a sintaxe básica da instrução for (sintaxe apresentada aqui).
  
	listaDeMercado := []string{"abacate", "sabonete", "azeite de oliva", "tomate", "banana", "macarrão", "cebola"}

	for indice := 0; indice < len(listaDeMercado); indice++ {
		fmt.Printf("%d) %s\n", indice+1, listaDeMercado[indice])
	}

}
