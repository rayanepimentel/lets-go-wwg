package main

import (
	"fmt"
)

func main() {
	for i := 13; i <= 27; i++{
		fmt.Println(i)
	}
	//para pegar o valor de i depois do for
	i := 13

	for i <= 27 {
		fmt.Println(i)
		i++
	}
	fmt.Println("fora do bloco do for ",i)//28

	hora := 0

	for hora < 24 {
		//fmt.Printf("%v:00\n", hora)
		fmt.Printf("%.2d:00\n", hora)
		hora++
	}

	for hora < 3 {
		for minuto := 0; minuto < 60; minuto++{
			for segundo := 0; segundo < 60; segundo++ {
				fmt.Printf("%.2d:%.2d:%.2d\n", hora, minuto, segundo)
			}
		}
		hora++
	}


}