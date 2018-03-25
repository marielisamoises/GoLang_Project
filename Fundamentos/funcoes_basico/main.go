package main

import (
	"fmt"

	"github.com/marielisamoises/cursoGo/GoLang_Project/Fundamentos/funcoes_basico/matematica"
)

func main() {
	resultado := matematica.Calculo(matematica.Multiplicador, 2, 2)
	fmt.Printf("o resultado da multiplicação é: %d\r\n", resultado)
	resultado = matematica.Soma(3, 3)
	fmt.Printf("o resultado da multiplicação é: %d\r\n", resultado)
}
