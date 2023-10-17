package main

import "fmt"

func main() {

	// PRIMEIRA FORMA DE CRIAR
	var var1 string = "Variável 1"

	// SEGUNDA FORMA DE CRIAR
	variavel2 := "Variável 2"

	fmt.Println(var1)
	fmt.Println(variavel2)

	// TERCEIRA FORMA DE CRIAR
	var (
		variavel3 string = "Variável 3"
		variavel4 string = "Variável 4"
	)

	fmt.Println(variavel3, variavel4)

	// QUARTA FORMA DE CRIAR
	variavel5, variavel6 := "Variável 5", "Variável 6"
	fmt.Println(variavel5, variavel6)

	// BASICAMENTE A CONSTANTE E A VARIAVEL SÃO DECLARADAS DA MESMA FORMA, MAS A CONSTANTE NÃO PODE TER O VALOR ALTERADO
	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	// TROCA DE VALORES
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)

}
