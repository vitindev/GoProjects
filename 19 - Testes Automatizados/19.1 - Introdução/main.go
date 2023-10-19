package main

import (
	"fmt"
	"testes-auto-introducao/enderecos"
)

/*
	Testes automatizados em Go, será uma função que vai
	testar outra função e ver se o resultado é o esperado
	Os testes são para garantir o resultado esperado

	O teste de uma função nunca fica no arquivo da função
	Porque os arquivos de testes, tem que ter um nome especifico

	Ferramentas de Cobertura de Testes:
	go test ./... - Ele fará os testes em todos os pacotes existentes no projeto
	go test -v - Modo verboso, onde vai escrever detalhes na tela
	go test --cover - Ele vai verificar a cobertura do nosso teste
	go test --coverprofile resultado.txt - Vai gerar uma log do que foi testado

	go tool cover --func=resultado.txt - Vai entender a log e jogar no terminal
	go tool cover --html=resultado.txt - Gera um arquivo HTML totalmente detalhado, com as linhas de código

*/

func main() {

	tipoEndereco := enderecos.TipoDeEndereco("Rodovia dos Bobos")
	fmt.Println(tipoEndereco)

}
