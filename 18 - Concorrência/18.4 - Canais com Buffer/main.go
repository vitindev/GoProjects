package main

import "fmt"

func main() {

	// Com Buffer no caso, especificamos a quantidade maxima do canal
	// Ele só vai bloquear quando tiver sua capacidade maxima
	canal := make(chan string, 2)
	canal <- "Olá Mundo!"
	canal <- "Programando em Go!"
	close(canal)

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

}
