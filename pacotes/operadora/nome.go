package operadora

import (
	"github.com/macrusal/cursodego/pacotes/prefixo"
	"strconv"
)
//NomeOperadora nome da operadora
var NomeOperadora = "VIVO"

//PrefixoDaCapitalOperadora prefixo e nome da operadora
var PrefixoDaCapitalOperadora = strconv.Itoa(prefixo.Capital) + " " + NomeOperadora

