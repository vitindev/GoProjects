package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint16
}

func main() {

	/* Primeiro forma de declarar*/

	var user usuario
	user.idade = 21
	user.nome = "Victor"
	user.endereco = endereco{"Joaquim Nabuco", 418}

	fmt.Println(user)

	/* Segunda forma de declarar*/

	user2 := usuario{"Victor", 21, endereco{"Joaquim Nabuco", 418}}
	fmt.Println(user2)

	/* Declarando somente um valor de uma struct*/

	user3 := usuario{idade: 21, endereco: endereco{numero: 35}}
	fmt.Println(user3)

}
