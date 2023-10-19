package main

import "fmt"

/* Passagem por parâmetro */
func inverterSinal(x int) int {
	return x * -1
}

/* Passagem por referência */
func inverterSinalComPonteiro(x *int) {
	*x = *x * -1
}

func main() {

	numero := 20
	fmt.Println(numero)
	inverterSinalComPonteiro(&numero)
	fmt.Println(numero)

}
