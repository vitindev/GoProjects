package main

import "fmt"

func main() {

	var numero int = 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else if numero == 15 {
		fmt.Println("igual a 15")
	} else {
		fmt.Println("Menor que 15")
	}

	// Criando uma variável e validando (if init)
	// Mas ela fica limitada somente ao scopo do conjunto do if

	if outroNum := numero; outroNum > 0 {
		fmt.Println("Número maior que 0")
	} else {
		fmt.Println("Número menor ou igual a 0")
	}

}
