package formas

import (
	"math"
	"testing"
)

/*
	Test Driven Development(TDD): Desenvolvimento guiado a testes
		Fazer o teste primeiro... isso é o que ele sugere
*/

func TestArea(t *testing.T) {

	t.Run("Retângulo", func(t *testing.T) {

		retangulo := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := retangulo.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A area recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}

	})

	t.Run("Circulo", func(t *testing.T) {

		circulo := Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circulo.area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A area recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}

	})

}
