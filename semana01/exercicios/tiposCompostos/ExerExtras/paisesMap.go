package main

import (
	"fmt"
)

func main() {
	cidadesPaises := map[string]string {
		"Fortaleza" : "Brasil",
		"Lima": "Peru",
		"Santiago": "Chile",
		"São Paulo": "Brasil",
		"Caracas": "Venezuela",
		"Santa Cruz de la Sierra": "Bolívia",
		"Barquisimeto": "Venezuela",
		"Montividéu": "Uruguai",
		"Buenos Aires": "Argentina",
	}
    
    frequencia := map[string]int {
		"Brasil" : 0,
		"Peru": 0,
		"Chile": 0,
		"Venezuela": 0,
		"Bolívia": 0,
		"Uruguai": 0,
		"Argentina": 0,
	}


    for k, _ := range frequencia {
        for _, v := range cidadesPaises {
            if k == v {
                frequencia[k]++
            }
        }
    }
    fmt.Println(frequencia)
}

