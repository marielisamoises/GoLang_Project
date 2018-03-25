package matematica

//Executa qualquer tipo de calculo
func Calculo(funcao func(int, int) int, numA int, numB int) int {
	return funcao(numA, numB)
}

//Multiplica um numero x * y
func Multiplicador(x int, y int) int {
	return x * y
}
