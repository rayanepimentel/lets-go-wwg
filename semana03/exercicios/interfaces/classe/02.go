package main

import (
	"fmt"
)

type Apresentando interface {
	Apresentar()
}

type Gato struct {
	nome string
	tipo string
}

func (g Gato) Apresentar() {
	fmt.Printf("Miauu, meu nome é %s e sou do tipo %s \n", g.nome, g.tipo)
}

type Passarinho struct {
	tipo string
	nome string
}

func (p Passarinho) Apresentar() {
	fmt.Printf("piu piu, meu nome é %s e sou do tipo %s \n", p.nome, p.tipo)
}


func main() {

	souGato := Gato{
		nome: "Olivia",
		tipo: "fofinho",
	}

	souPassarinho := Passarinho{
		nome: "Oscar",
		tipo: "tagarela",
	}

	demo := []Apresentando{souGato, souPassarinho}

	for _, valor := range demo {
		valor.Apresentar()
	}
}