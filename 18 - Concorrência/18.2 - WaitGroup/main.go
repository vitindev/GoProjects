package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	Paralelismo vs Concorrência
		O paralelismo acontece quando tem duas ou mais tarefas sendo executadas exatamente ao mesmo tempo.
		Tarefas que são executas concorrente não estão necessariamente sendo executadas ao mesmo tempo.
		(Uma não espera a outra terminar para rodar)
*/

/*
	GoRoutine: É uma função, metodo, que é invocado pela cláusula go
	Que é feito para não esperar uma operação terminar para executar a próxima
*/

func escrever(texto string) {

	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}

}

func main() {

	/* Método para sincronizar GoRoutines*/

	var waitGroup sync.WaitGroup
	waitGroup.Add(2) // Quantidade de GoRoutines

	go func() {
		escrever("Olá Mundo!")
		waitGroup.Done() // Terminou subtrair 1
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done() // Terminou subtrair 1
	}()

	waitGroup.Wait() // Esperar até que todas tenha terminado

}
