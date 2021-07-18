package main

import (
	"fmt"
)

func main() {
	soma := Soma(10, 5)

	fmt.Println(soma)
}

func Soma(a int, b int) int {
	return a + b
}