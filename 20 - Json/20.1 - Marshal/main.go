package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"` // Usamos isso para quando for convertido para json usar essa definição como chave
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

/*
Pacote para lidar com Json encoding/json
*/
func main() {

	//É feito para converter um map/struct para Json
	//json.Marshal()

	//É feito para converter um Json para um map/struct
	//json.Unmarshal()

	c := cachorro{"Rex", "Dalmata", 3}
	fmt.Println(c)

	sliceDoJson, erro := json.Marshal(c)

	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(bytes.NewBuffer(sliceDoJson)) //Retornar em formato json

	c2 := map[string]string{
		"nome": "Toby",
		"raca": "Poodle",
	}

	sliceDoJson2, erro := json.Marshal(c2)

	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(bytes.NewBuffer(sliceDoJson2)) //Retornar em formato json

}
