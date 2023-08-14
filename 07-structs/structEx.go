package main

import "fmt"

type Carro struct {
	Modelo     string
	Velocidade float64
	Cor        string
	Ligado     bool
}

func (c *Carro) ligar() {
	c.Ligado = !c.Ligado
}

func (c *Carro) acelerar(num float64) {
	c.Velocidade = c.Velocidade * num
}

func main() {
	ferrari := Carro{
		Modelo:     "Esportivo",
		Velocidade: 80.00,
		Cor:        "Vermelho",
		Ligado:     false,
	}

	fmt.Println(ferrari)

	ferrari.ligar()

	fmt.Println(ferrari)

	ferrari.acelerar(1.3)

	fmt.Println(ferrari)

}
