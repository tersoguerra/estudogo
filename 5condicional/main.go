package main

import "fmt"

func main() {
	// As condicionais são bem básicas, lembrando de não usar ()
	sentinela := true
	numero := 5

	if !sentinela {
		fmt.Printf("\nErro: Sentinela falseada\n")
	} else { // ELSE é obrigatoriamente na linha de fechar } do IF
		if numero >= 5 {
			fmt.Printf("\nBoa escolha de número\n")
		}

	}

	// Há um método de escrever linhas de código na condicional do IF e separar por ";", ter essa variável apenas no escopo do IF
	if numTexto, validado := descreverNumero(numero); validado {
		fmt.Printf("\nO número %s é alto.\n", numTexto)
	} else {
		numTexto := string(numero) // Declarar a mesma variável fora do IF comprova que ela só era válida no escopo do IF
		fmt.Printf("\nO número %s é baixo\n.", numTexto)
	}

	// Switch clássico, sem break
	mes := 1
	switch mes {
	case 2:
		fmt.Printf("\nMês de carnaval!!\n")
	case 1, 6, 12:
		fmt.Printf("\nMês de férias!!\n")
	default:
		fmt.Printf("\nMês chato...\n")
	}

	// Switch com condição nos cases
	switch {
	case sentinela && mes <= 12:
		fmt.Printf("\nO programa parece OK.\n")
	default:
		fmt.Printf("\nO programa ta esquisito!\n")
	}
}

func descreverNumero(numero int) (texto string, validador bool) {
	texto = string(numero)
	if numero >= 5 {
		validador = true
	}
	return
}
