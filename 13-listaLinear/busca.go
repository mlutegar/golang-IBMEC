package main

import (
	"fmt"
)

const (
	// MaxTam é o tamanho máximo da lista
	MaxTam = 10
)


func main() {
	lista := [MaxTam]int{2,1,17,9,12}

	fmt.Println(BuscaLinear1(lista, 5, 17))
	fmt.Println(BuscaLinear1(lista, 5, 9))

	fmt.Println(BuscaLinear2(lista, 5, 17))
	fmt.Println(BuscaLinear2(lista, 5, 9))
}

// BuscaLinear1 busca um valor em uma lista
func BuscaLinear1(lista [MaxTam]int, n, valor int) int {
	i := 0
	for i < n{
		if lista[i] == valor {
			return i
		}
		i++
	}
	return -1
}

// BuscaLinear2 busca um valor em uma lista
func BuscaLinear2(lista [MaxTam]int, n, valor int) int {
	i := 0
	lista[n] = valor
	for lista[i] != valor {
		i++
	}
	if i < n {
		return i
	}
	return -1
}