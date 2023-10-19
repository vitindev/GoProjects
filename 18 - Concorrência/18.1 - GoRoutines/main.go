package main

import (
	"fmt"
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

	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}

}

func main() {
	go escrever("Olá Mundo!")
	escrever("Programando em Go!")
}
