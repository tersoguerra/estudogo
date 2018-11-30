package matematica

/*
DivisaoComResto é uma função que retorna a divisão e o resto de 2 parâmetros inteiros como inteiros
É possível retornar 2 valores declarando 2 variáveis de retorno
*/
func DivisaoComResto(numero1 int, numero2 int) (resultado int, resto int) {
	resultado = numero1 / numero2
	resto = numero1 % numero2
	return
}
