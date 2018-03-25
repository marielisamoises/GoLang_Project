package main

import (
	"fmt"
)

type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func main() {
	casa := Imovel{}
	fmt.Printf("A casa é: %+v\r\n", casa)

	apartamento := Imovel{17, 56, "Home Sweet Home", 250000}
	fmt.Printf("O apartamento é: %v\r\n", apartamento)

	chacara := Imovel{
		Y:     85,
		Nome:  "Chacara",
		valor: 55,
		X:     22,
	}
	fmt.Printf("A chacara é: %v\r\n", chacara)

	casa.Nome = "Lar Doce Lar"
	casa.valor = 250000
	casa.X = 12
	casa.Y = 33
	fmt.Printf("A casa é: %v\r\n", casa)
}
