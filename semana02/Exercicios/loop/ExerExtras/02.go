package main

import (
	"fmt"
)

type Apt struct {
	numero              int
	nomeDaProprietaria  string
	possuiVagaDeGaragem bool
}

func main() {
	apt1 := Apt{
		numero:              001,
		nomeDaProprietaria:  "Marie",
		possuiVagaDeGaragem: true,
	}

	apt2 := Apt{
		numero:              002,
		nomeDaProprietaria:  "Olivia",
		possuiVagaDeGaragem: false,
	}

	apt3 := Apt{
		numero:              003,
		nomeDaProprietaria:  "Betts",
		possuiVagaDeGaragem: true,
	}

	apartamentos := []Apt{apt1, apt2, apt3}

	for _, apartamento := range apartamentos {
		possuiVaga := "Não"
		if apartamento.possuiVagaDeGaragem {
			possuiVaga = "Sim"
		}
		fmt.Printf("O apt número %d é da proprietária %s, e  ele tem vaga de caragem? %s\n", apartamento.numero, apartamento.nomeDaProprietaria, possuiVaga)
	}
}
