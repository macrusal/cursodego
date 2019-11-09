package matematica

//Calculo retorna qualquer tipo de calculo
func Calculo(funccao func(int, int)int, numeroA int, numeroB int) int  {

	return funccao(numeroA, numeroB)
}

//Multiplicador retorna multiplicacao entre dois numeros
func Multiplicador(x int, y int) int {
	return x * y
}
