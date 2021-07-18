package main

import (
	"fmt"
)

func main() {
	var idade int
	var hora int
	fmt.Println("Qual a sua idade? ")
	fmt.Scan(&idade)

	if idade < 18 {
		fmt.Println("Menor de idade")
	}else if (idade >= 18 && idade <= 60) {
		fmt.Println("Maior de idade")
	}else {
		fmt.Println("Idoso")
	}

	//switch
	fmt.Println("Que horas são? ")
	fmt.Scan(&hora)

	switch {
	case hora > 12 && hora < 18:
		fmt.Printf("Período de %f horas é tarde",hora)
	case hora <12 && hora > 5:
		fmt.Printf("Período de %f horas é manhã",hora)
	case hora <= 5 :
		fmt.Printf("Período de %f horas é madrugada",hora)
	case hora >= 18 && hora <= 24:
		fmt.Printf("Período de %f horas é noite",hora)
	default :
		fmt.Println("Informe uma hora  válida")
	}	


}