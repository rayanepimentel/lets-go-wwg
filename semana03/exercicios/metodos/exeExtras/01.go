package main

import (
	"fmt"
)

type stack []string

func (s stack) Push(v string) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, string) {
	tamanho := len(s)
	return s[:tamanho-1], s[tamanho-1]
}

func main() {
	s := make(stack, 0)
	s = s.Push("aqui")
	s = s.Push("ali")
	s = s.Push("acolá")

	fmt.Println(s)
	s, p := s.Pop()


	fmt.Printf("push: %v e pop: %v", s, p)

}
