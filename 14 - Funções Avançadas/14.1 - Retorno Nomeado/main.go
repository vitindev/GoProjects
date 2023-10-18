package main

import "fmt"

func calculosMatematicos(x, z int) (soma int, subtracao int) {
	soma = x + z
	subtracao = x - z
	return
}

func main() {

	soma, subtracao := calculosMatematicos(2, 3)
	fmt.Println("Soma=", soma, "Subtracao=", subtracao)

}
