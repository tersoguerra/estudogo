package matematica

/*
Calculadora é uma função que possuí como parâmetros uma outra função matemática para ser aplicada em 2 inteiros e os 2 inteiros que serão parâmetros dessa função
Uma função pode ser um parâmetro também ao colocar no lugar do <tipoParâmetro> os seguintes valores: func (<tipoParametroFunc>) <tipoRetornoFunc>
*/
func Calculadora(funcaoMatematica func(int, int) int, numero1 int, numero2 int) int {
	return funcaoMatematica(numero1, numero2)
}

// Multiplicacao é uma função que retorna a multiplicação de 2 parâmetros inteiros como inteiro
func Multiplicacao(numero1 int, numero2 int) int {
	return numero1 * numero2
}

// Divisao é uma função que retorna a divisão de 2 parâmetros inteiros como inteiro
func Divisao(numero1 int, numero2 int) (resultado int) {
	resultado = numero1 / numero2
	return
	/*
		Declarar no lugar <tipoRetorno> a seguinte notação: (<nomeVarRetorno> <tipoRetorno) , é possível declarar a variável que será retornada e apenas usar "return " que o compilador sabe o que deve ser retornado
		Obs: Ao declarar que o tipo da variárel é inteiro e executar uma função que pode retornar floar, como /, o compilador já converte pra inteiro
	*/
}
