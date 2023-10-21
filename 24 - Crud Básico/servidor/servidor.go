package servidor

import (
	"crud-basico/database"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    int32  `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

/*
Para comandos em GO sql
Usamos db.Prepare e para Consulta somente db.Query
*/

func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	corpo, erro := io.ReadAll(r.Body)

	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return //PODEMOS FAZER ISSO SÓ PARA PARAR A EXECUÇÃO
	}

	var usuario usuario

	if erro := json.Unmarshal(corpo, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	fmt.Println(usuario)

	db, erro := database.Conectar()

	if erro != nil {
		w.Write([]byte("Erro ao conectar ao Banco de Dados"))
		return
	}

	defer db.Close()

	statement, erro := db.Prepare("INSERT INTO usuarios (name, email) VALUES (?,?);")

	if erro != nil {
		w.Write([]byte("Erro ao criar um Prepared Statement"))
		return
	}

	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)

	if erro != nil {
		w.Write([]byte("Erro ao executar o statement"))
		return
	}

	idInserido, erro := insercao.LastInsertId()

	if erro != nil {
		w.Write([]byte("Erro ao obter o id inserido"))
		return
	}

	/*
		StatusCode
	*/
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! Id: %d", idInserido)))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {

	db, erro := database.Conectar()

	if erro != nil {
		w.Write([]byte("Erro ao conectar ao Banco de Dados"))
		return
	}

	defer db.Close()

	linhas, erro := db.Query("SELECT * FROM usuarios")

	if erro != nil {
		w.Write([]byte("Erro ao criar um Prepared Statement"))
		return
	}

	defer linhas.Close()

	var usuarios []usuario

	for linhas.Next() {

		var usuario usuario

		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao scanear o usuário"))
			continue
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)

	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Ocorreu um erro ao converter os usuários para JSON"))
		return
	}

}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)

	if erro != nil {
		w.Write([]byte("Ocorreu um erro ao converter o parâmetro para Inteiro"))
		return
	}

	db, erro := database.Conectar()

	if erro != nil {
		w.Write([]byte("Erro ao conectar ao Banco de Dados"))
		return
	}

	defer db.Close()

	linha, erro := db.Query("SELECT * FROM usuarios WHERE id=? LIMIT 1;", ID)

	if erro != nil {
		w.Write([]byte(fmt.Sprintf("Erro ao buscar usuario com id=%d", ID)))
		return
	}

	var usuario usuario

	if linha.Next() {

		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao scanear o usuário!"))
			return
		}

	}

	w.WriteHeader(http.StatusOK)

	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para JSON!"))
		return
	}

}
