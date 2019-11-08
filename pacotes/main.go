package main

import (
	"../pacotes/prefixo"
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
)


//NomeDoUsuario eh o nome do usuario do sisteme
var NomeDoUsuario = "Marcelo"

func main() {

	fmt.Printf("Nome do usuario  %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo da Capital: %d\r\n", prefixo.Capital)
}
