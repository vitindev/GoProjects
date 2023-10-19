package main

import "fmt"

/*
	Diferença entre método e função.
		Método: Está associado a algo (struct, interface)
*/

type usuario struct {
	nome  string
	idade uint8
}

/*
	Basicamente é uma função, porém está associada
	com um corpo de struct nesse caso usuario.
	Podemos passar parametros, retornos e afins
*/
func (user usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados\n", user.nome)
}

func (user usuario) maiorDeIdade() bool {
	return user.idade >= 18
}

/*
	Passando como ponteiro exatamente para conseguir modificar os campos da struct
*/
func (user *usuario) modificarIdade(novaIdade uint8) {
	user.idade = novaIdade
}

/* Função normal */
func escrever() {
	fmt.Println("Escrevendo...")
}

func main() {

	user := usuario{"Victor", 21}
	fmt.Println(user)
	user.salvar()
	user.modificarIdade(17)
	fmt.Println("É maior de idade:", user.maiorDeIdade())

}
