package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

/*
	Importar um pacote implicitamente: _ "github.com/go-sql-driver/mysql"
	Basta colocar o _ na frente do pacote

	A váriavel de error pode ser reutilizada

*/

func main() {

	urlConnection := "root:@/DevBook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", urlConnection)

	if erro != nil {
		log.Fatal(erro)
	}

	defer db.Close()

	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Conexão aberta com o Banco de Dados!")

	//PARA FAZER UMA QUERY

	linhas, erro := db.Query("SELECT * FROM usuarios;")

	if erro != nil {
		log.Fatal(erro)
	}

	defer linhas.Close()

	fmt.Println(linhas)

}
