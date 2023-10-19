package main

import "fmt"

/*
	Função executada antes da main
	Pode ter uma por arquivo, não uma por pacote
	Poder ser considerada uma forma de SETUP
*/
func init() {
	fmt.Println("Executando a função init")
}

func main() {
	fmt.Println("Executando a função main")
}
