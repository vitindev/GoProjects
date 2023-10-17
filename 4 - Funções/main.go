package main

import "fmt"

func main() {

	soma := somar(10, 20)
	fmt.Println(soma)

	// CRIANDO UMA FUNÇÃO COMO FORMA DE TIPO
	// AS FUNC podem ter mais de um retorno

	var funcao = func() {
		fmt.Println("Função F")
	}

	funcao()

	var funcao2 = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	funcao2("Olá amigos!")
	fmt.Println(funcao2("Hello World!"))

	resultadoSoma, resultadoSobtracao := calculosMatematicos(10, 20)
	fmt.Println(resultadoSoma, resultadoSobtracao)

	/*
		Se eu não quiser usar um resultado de uma função com vários retornos, posso ignorar com o _
	*/

	result, _ := calculosMatematicos(40, 50)
	fmt.Println(result)

}

/*
	Entendendo:
		Basicamente func para começar a função
		somar o nome
		x o nome da variavel int8 o tipo
		z o nome da variavel int8 o tipo
		int 8 no final o tipo de retorno da função

*/

func somar(x int8, z int8) int8 {
	return x + z
}

func calculosMatematicos(x, z int8) (int8, int8) {

	soma := x + z
	subtracao := x - z

	return soma, subtracao

}
