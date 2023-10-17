package main

import "fmt"

func main() {

	usuario := map[int]string{
		1: "Victor",
		2: "Jorge",
	}

	fmt.Println(usuario)

	// Para acessar o valor de uma chave, usando o tipo de chave referente ao map
	fmt.Println(usuario[1])

	// Aninhando MAPS

	usuario2 := map[int]map[string]string{

		1: {
			"Jorge":  "Mendes",
			"Marcos": "Vinicius",
		},
	}

	fmt.Println("before=", usuario2)

	// Para deletar chave do MAP, use a função nativa do go
	delete(usuario2, 1)

	// Para adicionar uma nova chave
	usuario2[2] = map[string]string{
		"Vila": "Velha",
		"Olá":  "Mundo",
	}

	fmt.Println("after=", usuario2)

}
