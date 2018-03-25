package matematica

//Executa qualquer tipo de calculo
func Calculo(funcao func(int, int) int, numA int, numB int) int {
	return funcao(numA, numB)
}

//Multiplica um numero x * y
func Multiplicador(x int, y int) int {
	return x * y
}

//Divisor efetua a divisão do numeroA pelo numeroB
func Divisor(numA int, numB int) (resultado int) {
	resultado = numA / numB
	return
}

//Resultado e resto
func DivisorComResto(numA int, numB int) (resultado int, resto int) {
	resultado = numA / numB
	resto = numA % numB
	return
}
