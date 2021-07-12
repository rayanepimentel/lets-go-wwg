// 1) Dado o ano em que a pessoa nasceu, calcule quantos anos ela tem no ano atual 
// (para fins de arredondamento, considere que ela já fez aniversário no ano atual).
// 2) Como podemos pegar a informação do ano diretamente do console?

package main

import(
	"fmt"
	"time"
)

func main()  {
	var birthDate int
	year := time.Now().Year()//pegando o ano atual
	fmt.Println("Qual o ano do seu nascimento ?")
	fmt.Scan(&birthDate)//salva a resposta em birtDate
	age := year - birthDate
	fmt.Printf("Sua idade é %v", age)
}