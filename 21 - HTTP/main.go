package main

import (
	"log"
	"net/http"
)

/*

	HTTP: é um protocolo de comunicação, base da comunicação WEB
		Funciona via cliente - servidor
			O cliente faz a requisição, o servidor processa e devolve uma resposta

	Tipos de mensagens:
		Request(Cliente) - Response(Servidor)

	Rotas: É feita para indentificar o tipo de mensagem que o cliente enviou localhost:8080/api/ exemplo de uma rota
		URI: Indentificar do recurso (Caminho)
		Método - O que vamos fazer com o URI que identifiquei (GET, POST, PUT, DELETE)
*/

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo!"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar páginas de usuários!"))
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {}) // RAIZ
	http.HandleFunc("/home", home)
	http.HandleFunc("/usuarios", usuarios)

	log.Fatal(http.ListenAndServe(":5000", nil))

}
