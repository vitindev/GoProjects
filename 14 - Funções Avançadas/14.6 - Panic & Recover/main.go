package main

import "fmt"

/*

	panic: Vai matar o programa!
		Antes de matar, vai chamar todas as funçãos com a clausula defer
		panic não é do tipo error

	recover: Recuperar um programa entrando em panic
		Não tem problema se usar o recover em uma função que não está em panic

*/

func calcularMedia(x, z float32) bool {

	defer recuperarExecucao()

	media := (x + z) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MÉDIA É EXATAMENTE 6!")
}

func recuperarExecucao() {

	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}

}

func main() {
	fmt.Println(calcularMedia(6, 6))
	fmt.Println("Pós execução!")
}
