package main

import "fmt"

var (
	//Endereco eh um valor importante
	Endereco string
	//Telefone eh um valor importante
	Telefone string
	quantidade, estoque int
	comprou bool
	valor float64
	palavras rune
)

func main() {
	teste := "Valor sem declarar"
	fmt.Printf("endereco: %s\r\n", Endereco)
	fmt.Printf("quantidade: %d\r\n", quantidade)
	fmt.Printf("comprou: %v\r\n", comprou)
	fmt.Printf("palavra: %v\r\n", palavras)
	fmt.Printf("teste: %s\r\n", teste)
}
