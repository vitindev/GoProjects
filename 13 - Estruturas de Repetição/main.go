package main

import (
	"fmt"
)

func main() {

	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando i = ", i)
	// 	time.Sleep(time.Second)
	// }

	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementando j = ", j)
	// 	time.Sleep(time.Second)
	// }

	numeros := []int{1, 2, 3}

	for indice, valor := range numeros {
		fmt.Println("Indice=", indice, "Valor=", valor)
	}

	//string(char) transforma os valores da tabela ASC para letras

	for indice, letra := range "PALAVRA" {
		fmt.Println("Indice=", indice, "Letra=", string(letra))
	}

	// Iterando um MAP
	// Não conseguimos iterar uma STRUCT

	usuario := map[int]string{
		1: "Nome",
		2: "Sobrenome",
	}

	for key, value := range usuario {
		fmt.Println("Key=", key, "Value=", string(value))
	}

	// LOOP INFINITO

	for {
		fmt.Println("Infinito...")
	}

}
