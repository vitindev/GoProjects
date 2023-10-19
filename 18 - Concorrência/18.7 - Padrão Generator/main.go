package main

import (
	"fmt"
	"time"
)

/*
	Padrão Generator: Basicamente ele vai ser uma função que vai
	encapsular uma chamada pra uma GoRoutine e devolver um canal.

	Ele é usado geralmente para esconder complexidade.

*/

func escrever(texto string) <-chan string {

	canal := make(chan string)

	go func() {

		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}

	}()

	return canal
}

func main() {

	canal := escrever("Olá mundo!")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}

}
