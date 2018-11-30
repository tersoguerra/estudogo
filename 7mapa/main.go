package main

type Exemplo struct {
	inteiro  int
	texto    string
	booleano bool
}

/*
Mapa é uma estutura que armazena um conjunto de dados e utiliza identificadores para localizar cada dado armazenado
A declaração do mapa segue: var <nome> map[<tipoIdentificador>] <tipoEstruturaArmazenada>
*/
var dicionario map[string]Exemplo

func main() {

	dicionario = make(map[string]Exemplo, 0)

	ex := Exemplo{11, "exemplo simples", true}
	nulo := Exemplo{}

}
