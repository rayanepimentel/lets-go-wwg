package main

import (
	"fmt"
)

type Pessoa struct {
	nome string
	idade int
	cor string
}

func main() {

	// Crie um mapa chamado ano onde as chaves (keys) são os números dos meses do ano (ex: 1, 2)
	// e os valores (values) são os nomes dos meses.
	// Printe os nomes do meses 1 e 12.
	// Printe o tamanho do mapa ano.

	//eu criei usando Map literal

	ano := map[int]string {
		1: "Janeiro",
		2: "Fevereiro",
		3: "Março",
		4: "Abril",
		5: "Maio",
		6: "Junho",
		7: "Julho",
		8: "Agosto",
		9: "Setembro",
		10: "Outubro",
		11: "Novembro",
		12: "Dezembro", 
	}

	fmt.Printf("%v, %v\n", ano[1], ano[12])
	fmt.Println(len(ano))

	//exer 07 usando struct 
	//Crie um tipo Pessoa que tenha os atributos nome, idade e cor preferida.
  //Crie três variáveis do tipo pessoa.
  //Printe o nome de todas as três, bem como frases informando sua idade e cores preferidas.

	pessoa1 := Pessoa {
		nome : "Talia",
		idade: 20,
		cor: "Azul",
	}

	pessoa2 := Pessoa {
		nome : "Ana",
		idade: 30,
		cor: "Lilás",
	}

	pessoa3 := Pessoa {
		nome : "Maria",
		idade: 40,
		cor: "Verde",
	}

	fmt.Printf("Nome das pessoas: %s, %s e %s\n", pessoa1.nome, pessoa2.nome, pessoa3.nome)
	fmt.Printf("A %s tem %d anos e sua cor preferida é %s\n", pessoa1.nome, pessoa1.idade, pessoa1.cor)
	fmt.Printf("A %s tem %d anos e sua cor preferida é %s\n", pessoa2.nome, pessoa2.idade, pessoa2.cor)
	fmt.Printf("A %s tem %d anos e sua cor preferida é %s\n", pessoa3.nome, pessoa3.idade, pessoa3.cor)



}