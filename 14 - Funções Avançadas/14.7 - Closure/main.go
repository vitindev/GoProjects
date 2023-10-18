package main

import "fmt"

/*
	Closure: São funções que referênciam variáveis que estão fora do seu corpo
*/

func closure() func() {

	texto := "Dentro da função Closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {

	texto := "Dentro da função main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()

}
