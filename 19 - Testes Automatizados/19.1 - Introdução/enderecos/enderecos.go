package enderecos

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func TipoDeEndereco(endereco string) string {

	tiposValidos := []string{"rua", "avenida", "rodovia"}
	primeiraPalavra := strings.Split(strings.ToLower(endereco), " ")[0]

	for _, tipo := range tiposValidos {

		if tipo == primeiraPalavra {
			return cases.Title(language.BrazilianPortuguese).String(tipo)
		}

	}

	return "Não tem um tipo válido!"
}
