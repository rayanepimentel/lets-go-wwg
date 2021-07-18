package main

import (
	"fmt"
)

var nome = "Olivia"
var hora = 11

func main(){
	cumprimento := GereCumprimente(nome, hora)
	fmt.Println(cumprimento)
}

func GereCumprimente(nome string, hora int) string{

	cumprimento := ""
	switch  {
	case hora >= 6 && hora < 12:
		cumprimento = "Bom diaa"
	case hora >= 12 && hora < 18:
		cumprimento = "Boa Tarde"
	case hora >= 18 && hora < 24:
		cumprimento = "Boa noitee"
	case hora >= 0 && hora <6:
		cumprimento = "Boa madrugada"
	default:
		cumprimento = "Olar"
	}
	frase := fmt.Sprintf("%s, %s !", cumprimento, nome)
	return frase
}