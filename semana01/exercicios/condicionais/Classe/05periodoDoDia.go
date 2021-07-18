package main

import (
	"fmt"
)

func main() {
	peridoDoDia := 13

	switch {
	case peridoDoDia >= 6 && peridoDoDia < 12:
		fmt.Println("é manhã")
	case peridoDoDia >= 12 && peridoDoDia < 18:
		fmt.Println("é tarde")
	case peridoDoDia >= 18 && peridoDoDia < 24:
		fmt.Println("é noite")
	case peridoDoDia >= 0 && peridoDoDia < 6:
		fmt.Println("é madrugada")
	default:
		fmt.Println("Ops essa não é uma hora válida")
	}

}