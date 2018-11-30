package model /* É uma boa prática utilizar  o pacote model para isolar as estruturas */

import (
	"encoding/json" // Biblioteca nativa json
)

// Celular é uma estrutura pública que possui campos comuns para celulares, como públicos modelo e número e privado valor
type Celular struct {
	Modelo string `json:"modelo"` // `` é utilizado para fazer notações na estrutura e essa notação em específico indica para a biblioteca nativa json que deve user esse nome para o campo do JSON
	Numero int    `json:"numero"`
	valor  float64
}

/*
A forma de atribuir funções à uma estrutura é criar uma função no pacote da estrutura sinalizando antes do nome da estrutura:
(<nomeInstancia> *<nomeEstrutura>)
*/

//GetValor retorna o valor de um Celular
func (c *Celular) GetValor() float64 {
	return c.valor
}

//SetValor altera o valor de um Celular
func (c *Celular) SetValor(valor float64) {
	c.valor = valor
}

//ExportJSON retorna uma string do objeto Celular na notação de JSON
func (c *Celular) ExportJSON() string {
	objJSON, _ := json.Marshal(c) // _ é uma variável padrão para indicar que não será utilizado, pois é retornado uma sentinela de erro
	return string(objJSON)
}
