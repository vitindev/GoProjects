package main

import (
	"errors"
	"fmt"
)

func main() {

	/*
		Tipos Númericos Inteiro:

			Basicamente o que vem na frente como 8, 16, 32 e 64 são
			a quantidade de bits que eles conseguem armazenar

			int8, int16, int32, int64
			int - Vai usar a arquitetura do meu computador, Win64x = int64
			uint - UNSIGNED int, que basicamente não possui sinal.

			Alias:
			var variavel rune(int32) = 12456
			var variavel byte(int8) = 123
			byte = 8 bits, por isso o alias

	*/

	var numero int = -100000
	fmt.Println(numero)

	var unsigned uint = 100000
	fmt.Println(unsigned)

	/*
		Tipos Númericos Reais:

			Basicamente o que vem na frente como 32 e 64 são
			a quantidade de bits que eles conseguem armazenar

			Podemos usar a inferencia de tipo := porém o retorno é float
			mas não podemos criar uma variavel usando float como -> var real float

	*/

	var real1 float32 = 123.45
	fmt.Println(real1)

	var real2 float64 = 1233333.45
	fmt.Println(real2)

	/*

		Tipos de String:

			Só existe a string, não existe nem char.
			Caso use implicitamente o aspas simples(char), ele vai retornar um número.

	*/

	var str string = "String 1"
	fmt.Println(str)

	str2 := "String 2"
	fmt.Println(str2)

	/*

		Valor 0:
			É quando a gente criar uma variável, mas não atribuir um valor.
			Então basicamente cada tipo de dado tem seu valor 0, como nil, 0, "", false

	*/

	var texto string
	fmt.Println(texto)

	/*

		Tipo Booleano:
			bool
			Basicamente true or false

	*/

	var booleano bool = true
	fmt.Println(booleano)

	var boolean bool = false
	fmt.Println(boolean)

	/*

		Tipo erro:
			error
			Basicamente é o que usamos para tratar erros e afins
			Para criar um error precisamos usar uma dependencia nativa do go
			"errors"

	*/

	var erro error = errors.New("Erro Interno")
	fmt.Println(erro)

}
