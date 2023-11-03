/*
	Michel Lutegar D'Orsi Pereira
	14/08/2023
*/

package main

import (
	"fmt"
	"math"
)

// É parecido com uma classe, mas não é. Não é POO. É uma forma de agrupar elementos pertecentes ao mesmo contexto.

type Pessoa struct {
	Nome   string
	Idade  int
	Altura float64
}

type Ponto struct {
	X, Y int
}

type Retangulo struct {
	Ponto
	Largura, Altura int
}

type Circulo struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

func main() {
	pessoa1 := Pessoa{
		Nome:   "Alice",
		Idade:  25,
		Altura: 2.2,
	}

	fmt.Println(pessoa1)
	fmt.Println(pessoa1.Nome)
	fmt.Println(pessoa1.Altura)

	retangulo := Retangulo{
		Ponto:   Ponto{X: 1, Y: 2},
		Largura: 60,
		Altura:  34,
	}

	fmt.Println(retangulo)

	circ := Circulo{
		Raio: 2.3,
	}

	fmt.Println(circ.Area())
}
