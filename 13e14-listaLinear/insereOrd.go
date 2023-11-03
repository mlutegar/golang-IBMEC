/*
	Michel Lutegar D'Orsi Pereira
	24/09/2023
*/

package main

import "fmt"

const M = 10

func main() {
	var v [M]int = [M]int{1, 3, 4, 5, 6, 7, 8, 9}
	var n int = 8
	insereOrd(&v, &n, 2)
	fmt.Println(v)
}

func insereOrd(v *[M]int, n *int, novoValor int) {
	if *n == M {
		fmt.Println("Overflow")
		return
	}

	fmt.Println("Tentando inserir", novoValor)
	aux := novoValor
	for i := 0; i < *n; i++ {
		if v[i] == novoValor {
			fmt.Println("Valor já existe")
			return
		}
		if v[i] > novoValor {
			v[i], aux = aux, v[i]
		}
	}
	fmt.Println(novoValor, "não encontrado, inserindo na lista")
	v[*n] = aux
	*n++
}
