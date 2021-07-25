package main

import (
	"fmt"
	"strconv"
)

// CPF simula o documento de identificação de pessoa física brasileira
type CPF string


func main() {
	cpf := CPF("1111111130")
	if cpf.EhValido() {
		fmt.Println("CPF válido")
	} else {
		fmt.Println("CPF inválido")
	}
}

// EhValido retorna se um CPF é valido ou não
func (c CPF) EhValido() bool {
	_, err := strconv.Atoi(string(c))
	if err == nil && len(c) == 11  {
		return true
	}
	return false
}