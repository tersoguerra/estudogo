// Comentários são iniciados com: //
// Curso: udemy.com/cursodego/
// Obs1: Caracteres UTF-8 Unicode por padrão
// Obs2: Auto-identação

package main // Declarando como pacote básico

import "fmt" // Importando pacote com funções básicas

var variavel string               // Declaração de variável: var <nome> <tipo>
var texto1 = "Declaração simples" // Declaração pode ser feita atribuindo valor, definindo tipo: var <nome> = <valor>

// Variáveis são inicalizadas com valor padrão por tipo
var initString string // initString = “”
var initInt int       // initInt = 0
var initFloat float64 //  initFloat = 0.00
var initBool bool     // initBool = false

// Publico deve ter uma linha de comentário para explicar a necessidade de ser público
var Publico string // Variáveis públicas apenas necessitam iniciar com letra maiúscula

func main() { // Função básica de inicialização, famoso main()
	// Valores de texto 1 e 2 não podem ser utilizados nessa função, pois seus valores estão fora do escopo desse bloco
	texto3 := "Valor qualquer" // Declaração sem tipo de variável local pode utilizar ":=" sem "var"
	// Variáveis declaradas sem utilização fazem o programa NÃO COMPILAR.
	// fmt.PrintF() output de string, utilizando função (<string>, <valor1>, <valor2>)
	fmt.Printf("texto = %v\n", texto3)          // %v insere qualquer valor do parâmetro após a string no local indicado
	fmt.Printf("initString = %s\n", initString) // %s: string
	fmt.Printf("initInt = %d\n", initInt)       // %d: int
	fmt.Printf("initFloat = %f\n", initFloat)   // %f: float
	fmt.Printf("initBool = %t\n", initBool)     // %t: bool
}
