package main

import (
	"fmt"
	"pacotes/auxiliar"
)

/*

	go mod tidy - Remove dependencias externas que não estão sendo utilizadas.
	go get (url) - Adiciona dependencias externas no go.mod para serem utilizados.
	go build - Compila a aplicação.
	go run (arquivo.go) - Iniciar uma aplicação go
	go mod init (nome do mod)

*/

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()
}
