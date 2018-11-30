package main //Daclaração de que esse é o pacote principal do sistema, indicando que método main() estará aqui e será o início do sistema

import "fmt"     //Importação de pacote do Go com funções básicas
import "./anexo" //Importação do meu pacote (anexo) que está abaixo (./) do pacote atual na hierarquia de arquivos

func main() {
	//Para importar variáveis públias de outro arquivo, declarar pacote no novo, declarar variável e importar em qqr arquivo: <pacote>.<variávelPública>
	fmt.Printf("Enviando valor de variável de outro arquivo no pacote 'anexo': %s\n", anexo.VariavelAnexa)
}
