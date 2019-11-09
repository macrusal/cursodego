package main

import (
	"fmt"
	"github.com/macrusal/cursodego/funcoes/basico/matematica"
)

func main()  {
	resultado := matematica.Calculo(matematica.Multiplicador,2, 2)
	fmt.Printf("O resultado da multiplicação é: %d\r\n", resultado)
	resultado = matematica.Soma(resultado, 25)
	fmt.Printf("O resultado da soma é: %d\r\n", resultado)
}
