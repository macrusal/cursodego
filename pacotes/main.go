package main

import (
	"fmt"
	"github.com/macrusal/cursodego/pacotes/operadora"
	"github.com/macrusal/cursodego/pacotes/prefixo"
)

//NomeDoUsuario eh o nome do usuario do sisteme
var NomeDoUsuario = "Marcelo"

func main() {
	fmt.Printf("Nome do usuario  %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo da Capital: %d\r\n", prefixo.Capital)
	fmt.Printf("Operadora: %s\r\n",  operadora.NomeOperadora)
	fmt.Printf("Teste: %s\r\n", prefixo.CapitalNomeCapital)
}
