package main

import "fmt"

func fibonnaci(posicao int) int {

	if posicao <= 1 {
		return posicao
	}

	return fibonnaci(posicao-2) + fibonnaci(posicao-1)
}

// tarefas com o parametro <-chan só recebe dados
// resultados com o parametro chan<- só envia dados
func worker(tarefas <-chan int, resultados chan<- int) {

	for numero := range tarefas {
		resultados <- fibonnaci(numero)
	}

}

func main() {

	tarefas, resultados := make(chan int, 45), make(chan int, 45)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}

	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}

}
