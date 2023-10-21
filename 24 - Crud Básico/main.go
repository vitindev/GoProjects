package main

import (
	"crud-basico/servidor"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*
	CRUD = CREATE, READ, UPDATE, DELETE

		CREATE - POST
		READ - GET
		UPDATE - PUT
		DELETE - DELETE

*/

func main() {

	router := mux.NewRouter() // Vai conter todas as rotas
	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost)
	router.HandleFunc("/usuarios", servidor.BuscarUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", servidor.BuscarUsuario).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":5000", router))

}
