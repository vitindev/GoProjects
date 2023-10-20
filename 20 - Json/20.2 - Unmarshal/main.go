package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {

	//`json:"-"` ignora o nome do struct

	Nome  string `json:"nome"` // Usamos isso para quando for convertido para json usar essa definição como chave
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

/*
Pacote para lidar com Json encoding/json
*/
func main() {

	cachorroEmJson := `{"nome":"Rex","raca":"Dalmata","idade":3}`

	var c cachorro

	if erro := json.Unmarshal([]byte(cachorroEmJson), &c); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c)

	cachorroEmJson2 := `{"nome":"Toby","raca":"Poodle"}`

	c2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(cachorroEmJson2), &c2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c2)

}
