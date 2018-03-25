package operadora

import (
	"strconv"

	"github.com/marielisamoises/cursoGoLang/variaveis/pacotes/prefixo"
)

//NomeOperadora representa o nome que se estar a usar
var NomeOperadora = "XPOT Telecom"

var PrefixoDaCapitalOperadora = strconv.Itoa(prefixo.Capital) + " " + NomeOperadora
