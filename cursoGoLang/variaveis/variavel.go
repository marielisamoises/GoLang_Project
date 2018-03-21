package main

import (
	"fmt"
)

var (
	Nome, telefone string  //Nome = variavel publica, telefone = variavel privada
	quantidade     int     //quantidade = 0
	comprou        bool    //comprou = false
	valor          float64 // valor = 0.00
	palavras       rune
)

func main() {

	endereco := "teste de variavel"
	fmt.Printf("endereco: %s\r\n", endereco)
	fmt.Printf("quantidade: %d\r\n", quantidade)
	fmt.Printf("telefone: %s\r\n", telefone)
	fmt.Printf("comprou: %v\r\n", comprou)
	fmt.Printf("palavras: %v\r\n", palavras)
}
