package main

import "fmt"

func main() {

	/* Ponteiro é uma referência de memória */
	/* Basicamente a gente faz a referência para o endereço onde está alocado aquele atributo */

	var variavel int = 10
	var variavel2 int = variavel

	fmt.Println(variavel, variavel2)
	variavel++

	/* Aplicando ponteiros */

	var variavel3 int = 100
	var ponteiro *int = &variavel3

	fmt.Println(variavel3, *ponteiro)

}
