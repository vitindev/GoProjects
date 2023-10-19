package main

import (
	"fmt"
	"time"
)

/*
	Canal: É porque é um canal de comunicação,
	que pode tanto enviar como receber dados,
	e é com isso que vamos sincronizar as GoRoutines

	O canal é definido por chan

	O canal tem duas operações, enviar e receber
	Elas bloqueiam a execução do programa

	O canal é uma operação bloqueante

*/

func escrever(texto string, canal chan string) {

	for i := 0; i < 5; i++ {

		//Enviando um valor para o canal
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal)

}

func main() {

	canal := make(chan string)

	go escrever("Olá Mundo!", canal)

	// for {

	// 	//Recebendo um valor, esperando...
	// 	mensagem, aberto := <-canal

	// 	//Verificando se o canal está aberto
	// 	if !aberto {
	// 		break
	// 	}

	// 	fmt.Println(mensagem)
	// }

	// Mesma sintaxe acima, porém melhorada
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do Programa!")

}
