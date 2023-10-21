package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Conectar() (*sql.DB, error) {

	url := "root:@/DevBook?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", url)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
