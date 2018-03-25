package main

import (
	"fmt"

	"github.com/marielisamoises/cursoGoLang/variaveis/pacotes/operadora"
	"github.com/marielisamoises/cursoGoLang/variaveis/pacotes/prefixo"
)

//NomeDoUsuario é o nome do usuário no sistema
var NomeDoUsuario = "Mari"

func main() {
	fmt.Printf("Nome do usuário: %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo do DDD: %d\r\n", prefixo.Capital)
	fmt.Printf("Nome da Operadora: %s\r\n", operadora.NomeOperadora)
}
