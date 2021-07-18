
package main

import (
	"fmt"
	"time"
)

func main() {
	//Declare uma variável e atribua a ela o ano de nascimento de uma pessoa.
	var anoDeNascimento int
	//Declare uma variável e atribua a ela o ano atual.
	anoAtual := time.Now().Year()
	fmt.Println("Qual o ano do seu nascimento ?")
	fmt.Scan(&anoDeNascimento)//salva a resposta em birtDate
	idade := anoAtual - anoDeNascimento

	if idade >= 16 {
		fmt.Println("Olá, você já tá apata a votar!")
		return
	}
	fmt.Println("Desculpe, você ainda não tem idade para votar :(")


}