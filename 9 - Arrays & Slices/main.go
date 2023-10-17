package main

import "fmt"

func main() {

	/*
		Array:
			Lista de valores, uma variavel que guarda vários elementos.
			Ele possue tamanho fixo
	*/

	// Declarando o array modo 1

	var array [5]int
	array[0] = 1
	fmt.Println(array)

	// Declarando o array modo 2

	array2 := [5]float32{3, 5, 2.5, 1, 9.0}
	fmt.Println(array2)

	// Os 3 pontos ali dentro, vão se adequar a quantos elementos eu alocar
	// Não é DINÂMICO

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	/*

		Slice:
			Baseado em array só que dinâmico

		append do Slice, pega o slice atual e adiciona mais uma posição com um determinado valor

	*/

	slice := []int{10, 11, 13, 14, 15, 15}
	fmt.Println(slice)

	slice = append(slice, 18)
	fmt.Println(slice)

	/*
		Pega um intervalo da array e coloca em um slice
		Basicamente fazendo um papel de ponteiro
		Slice aponta para um array
	*/

	slice2 := array2[1:3]
	fmt.Println(slice2)

	// Alterando a posição do array e quando printa, mostra que atualizou também no slice

	array2[1] = 8
	fmt.Println(slice2)

	/*
		Arrays Internos
			Função make, vai alocar um espaço na memória
			Função len mostra a quantidade de itens que tem no slice/array
			Função cap mostrar a capacidade de itens que o slice/array suporta
	*/

	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // Capacidade

	/*
		Basicamente quando estoura a capacidade, ele pega o length e multiplica por 2,
		exatamente para não fica sem espaço
	*/

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println("Slice4=", len(slice4)) // length
	fmt.Println("Slice4=", cap(slice4)) // Capacidade

}
