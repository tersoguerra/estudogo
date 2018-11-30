package main

import (
	"fmt"

	"./matematica"
)

func main() {
	numero1 := 9
	numero2 := 4
	resultado := matematica.Soma(numero1, numero2)
	fmt.Printf("\nO resultado de aplicar matematica.Soma em %d e %d = %d\n", numero1, numero2, resultado)
	resultado = matematica.Calculadora(matematica.Multiplicacao, numero1, numero2)
	fmt.Printf("\nO resultado de aplicar matematica.Calculadora em matematica.Multiplicacao, %d e %d = %d\n", numero1, numero2, resultado)
	resultado = matematica.Calculadora(matematica.Divisao, numero1, numero2)
	fmt.Printf("\nO resultado de aplicar matematica.Calculadora em matematica.Divisao, %d e %d = %d\n", numero1, numero2, resultado)
	resultado, resto := matematica.DivisaoComResto(numero1, numero2)
	fmt.Printf("\nO resultado de aplicar matematica.DivisaoComResto em %d e %d = %d e resto %d\n", numero1, numero2, resultado, resto)
}
