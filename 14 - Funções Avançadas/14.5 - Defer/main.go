package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func funcao3(x, z float32) bool {

	fmt.Println("Verificando na função para aprovação do aluno")

	result := (x + z) / 2

	fmt.Println("Média calculada. Resultado será retornado!")

	if result >= 6 {
		return true
	}

	return false
}

func main() {

	// DEFER = Adiar, até o último momento possível
	// Geralmente usado com o banco de dados

	defer funcao1()
	funcao2()

	fmt.Println(funcao3(5, 7))

}
