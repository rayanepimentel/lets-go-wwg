package main 

import (
	"fmt"
)

func main() {

	var timeVermelho = []string{"Helena", "Jonas", "José", "Juliana"}

	timeVermelho = append(timeVermelho, "Luiz Inácio")

	fmt.Println(timeVermelho)
}