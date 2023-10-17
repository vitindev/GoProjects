package main

import "fmt"

type pessoa struct {
	nome  string
	idade uint8
}

type estudante struct {
	pessoa    // Utilizando assim estamos copiando os tipos de pessoa e colocando em estudante
	curso     string
	faculdade string
}

func main() {

	/* Em sí o GO não tem herença ou polimorfismo */

	var student estudante
	student.nome = "Victor"
	student.idade = 21
	student.curso = "Ciencia da Computação"
	student.faculdade = "Universidade Federal de Jataí"

	fmt.Println(student)

}
