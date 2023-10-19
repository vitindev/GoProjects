package main

import (
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {

	/*
	   os.Args: Reconhece o que é passado pela linha de comando via terminal ou afins
	*/

	/*
		Para iniciar o app:

		Ação 1: Listar os IPs
		go run main.go ip --host amazon.com.br

		Ação 2: Listas os Servidores
		go run main.go servidores --host amazon.com.br
	*/

	aplicacao := app.Gerar()

	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}

}
