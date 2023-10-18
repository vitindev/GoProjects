package main

import "fmt"

func main() {

	/* Cria a função e quando está pronta já executa */

	func() {
		fmt.Println("Hello World!")
	}()

	/* Função com parâmetro */

	func(texto string) {
		fmt.Println(texto)
	}("Parametro")

	/* Função com retorno */

	retorno := func(x, z int) int {
		return x + z
	}(3, 2)

	fmt.Println(retorno)

}
