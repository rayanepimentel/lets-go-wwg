package main

import (
	"fmt"
)

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
}