package main

import (
	"fmt"

	"github.com/marielisamoises/cursoGo/GoLang_Project/Fundamentos/funcoes_avancado/matematica"
)

func main() {
	resultado := matematica.Calculo(matematica.Multiplicador, 2, 2)
	fmt.Printf("o resultado da multiplicação é: %d\r\n", resultado)
	resultado = matematica.Soma(3, 3)
	fmt.Printf("o resultado da soma é: %d\r\n", resultado)
	resultado = matematica.Calculo(matematica.Divisor, 12, 3)
	fmt.Printf("o resultado da divisão é: %d\r\n", resultado)
	resultado, resto := matematica.DivisorComResto(15, 2)
	fmt.Printf("o resultado da divisão foi: %d e o resto da divisão foi %d\r\n ", resultado, resto)
}
