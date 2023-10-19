package enderecos

import "testing"

/*
	Teste Unitário

	Começa obrigatoriamente com a palavra
	Test + NomeDaFunção(Não precisa ser o mesmo nome, mas é bom por ser uma boa prática)

	O arquivo precisa ter o _test no nome do arquivo
*/
func TestTipoDeEndereco(t *testing.T) {

	t.Parallel() // Fala que pode rodar em paralelo

	endereco := "Rua ABC"
	tipoDeEnderecoEsperado := "Rua"
	tipoDeEnderecoRecebido := TipoDeEndereco(endereco)

	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
		t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s", tipoDeEnderecoEsperado, tipoDeEnderecoRecebido)
	}

}

/*
	Teste Unitário com Diversos Cenários
*/

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco2(t *testing.T) {

	t.Parallel() // Fala que pode rodar em paralelo

	cenarios := []cenarioDeTeste{
		{"RUA DOS BOBOS", "Rua"},
		{"Rua ABC", "Rua"},
		{"Avenida Brasil", "Avenida"},
		{"avenida FRANCO costa", "Avenida"},
		{"Rodovia Franco", "Rodovia"},
		{"Praça da Sé", "Não tem um tipo válido!"},
		{"praça da rodovia", "Não tem um tipo válido!"},
		{"", "Não tem um tipo válido!"},
	}

	for _, cenario := range cenarios {

		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		tipoDeEnderecoEsperado := cenario.retornoEsperado

		if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s", tipoDeEnderecoRecebido, tipoDeEnderecoEsperado)
		}

	}

}
