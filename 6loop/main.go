package main

import (
	"fmt"
)

func main() {
	/*Em Go, só existe "for" como palavra reservada de loop*/

	// for padrão = for com iterador declarado na condicional separado por ";", apenas para o escopo do for
	for i := 0; i <= 5; i++ {
		fmt.Printf("Valor de i é: %d\n", i)
	}

	// while = for com condifional
	i := 5
	sentinela := true
	for sentinela {
		i--
		if i <= 0 {
			sentinela = false
		}
		fmt.Printf("Valor de i é: %d\n", i)
	}

	// while True = não ter conficional
	for {
		i++
		fmt.Printf("Valor de i é: %d\n", i)
		if i >= 10 {
			break
		}
	}

	/*
		for in range = iterar em lista, retornando índice e valor
		Utilizar: for <variávelÍndice>, <variávelItem> := range <lista>
	*/
	texto := "lol Louco Loop"
	for indice, letra := range texto {
		fmt.Printf("texto[%d] = %c\n", indice, letra)
	}
}
