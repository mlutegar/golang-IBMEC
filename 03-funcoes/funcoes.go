/*
	Michel Lutegar D'Orsi Pereira
	09/08/2023
	ANOTADO
*/

package main

import "fmt"

func main() {
	fmt.Println(soma(23, 6))
	fmt.Println(subtracao(2, 5))

	var x, y = 2.4, 4.6
	x, y = trocaValores(x, y)
	fmt.Println(x, y)

	anonima()
}

// func nome(param1 tipo, param2 tipo, ...) tipoRetorno{}
func soma(a int, b int) int {
	return a + b
}

// Nesse caso, em que os dois parâmetros são do mesmo tipo,
// não há necessidade de declarar o tipo duas vezes
func subtracao(a, b float64) float64 {
	return a - b
}

// É possível que uma função retorne mais de um valor
func trocaValores(a, b float64) (float64, float64) {
	return b, a
}

// Função anônima (closure)
// É bom utilizar funções anonimas para funções privadas,
// que será utilizada apenas dentro do escopo de uma função
//"" especifica

func anonima() {
	dobra := func(x int) int {
		return x * 2
	}

	resultado := dobra(5)
	fmt.Println(resultado)

	calcular := func(f func(int) int, x int) int {
		return f(x)
	}

	resultado = calcular(dobra, 8)
	fmt.Println(resultado)
}

// Um atalho interessante é o de ir para o console e para o bloco de código. Use
// o Ctrl + ', para fazer isso

// Vê como funciona o lambda dentro do python
