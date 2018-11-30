package main

import (
	"fmt"

	"./model"
)

/*
Pessoa é uma estrutura pública que possuí campos comuns para pessoas, como públicos nome e idade e privado peso
Estrutura de dados são criadas seguindo: type <nome> struct{<nomeCampo> <tipoCampo>}
*/
type Pessoa struct {
	Nome  string
	Idade int
	peso  float64
}

func main() {
	isentao := Pessoa{}                         // Uma estrutura é declarada vazia como: <nomeEstrutura>{}
	fmt.Printf("\nO isentao é: %+v\n", isentao) // Para atribuir em uma string o valor da estrutura com identificação dos campos, utilizar: %+v
	usuario := Pessoa{"Cleber", 35, 80.7}       // É possível declarar uma estrutura indicando os valores na ordem da declaração, separados por virgula
	fmt.Printf("\nO usuário é: %+v\n", usuario)
	admin := Pessoa{ // É possível também declarar uma estrutura da seguinte forma: <nomeEstrutura>{<nomeCampo>: <valorCampo>,}
		Idade: 21,
		Nome:  "Administrador",
		peso:  62.4,
	}
	fmt.Printf("\nO admin é: %+v\n", admin)
	isentao.Nome = "Isento" // Para mudar o valor de um campo, atribuir o valor para <nomeVariavel>.<nomeCampo>
	fmt.Printf("\nO isentao é: %+v\n", isentao)
	// Ponteiros = endereço da memória que aponta para onde o dado está armazenado
	ponteiroPessoa := new(Pessoa) // Declarar estrutura com new(<nomeEstrutura>) retorna o ponteiro da estrutura
	fmt.Printf("\nO ponteiro de Pessoa indica:\n")
	fmt.Printf("Endereço: %p\n", &ponteiroPessoa) // O sinalizador de ponteiro é %p e para retornar o endereço do ponteiro, usar
	fmt.Printf("Conteúdo: %+v\n", ponteiroPessoa)
	esvaziarPessoa(&isentao)
	fmt.Printf("\nNovo valor do insentao ao ser esvaziado por referência:\n%+v\n", isentao)
	// Estrutura Avançada do pacote Model
	celular := model.Celular{
		Modelo: "Iphone",
		Numero: 987654321,
	}
	celular.SetValor(30150.99)
	fmt.Printf("\nEstrutura Celular exportada como JSON:\n%s\n", celular.ExportJSON())
	fmt.Printf("E atributo privado 'valor' é: %f\n", celular.GetValor())
}

/*
Por padrão, os parâmetros sempre são recebidos por valor, ao invés de por referência
Passar parâmetro como ponteiro para:
1- poder alterar uma variável fora do seu escopo
2- poder economizar memória, passando um endereço ao invés de toda estrutura
*/
func esvaziarPessoa(pessoa *Pessoa) { // Para sinalizar ponteiro, usasr: *<tipoParâmetro>
	pessoa.Nome = ""
	pessoa.Idade = 0
	pessoa.peso = 0.0
}
