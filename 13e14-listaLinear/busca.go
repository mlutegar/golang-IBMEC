package main

import (
	"fmt"
)

const (
	// MaxTam é o tamanho máximo da lista
	MaxTam = 10
)


func main() {
	lista := [MaxTam]int{2, 5, 7, 9, 12, 14}
	var tam int = 6

	// BuscaLinear1
	fmt.Println("\nBuscaLinear1")
	fmt.Println(BuscaLinear1(lista, MaxTam-1, 17))
	fmt.Println(BuscaLinear1(lista, MaxTam-1, 9))

	// BuscaLinear2
	fmt.Println("\nBuscaLinear2")
	fmt.Println(BuscaLinear2(lista, MaxTam-1, 17))
	fmt.Println(BuscaLinear2(lista, MaxTam-1, 9))

	// buscaBinaria
	fmt.Println("\nbuscaBinaria")
	fmt.Println(buscaBinaria1(lista, MaxTam, 17))
	fmt.Println(buscaBinaria1(lista, MaxTam, 9))

	// insere item novo
	fmt.Println("\nInsere item novo")
	Insere(&lista, &tam, 17)
	fmt.Println(lista)

	// insere item já existente
	fmt.Println("\nInsere item já existente")
	Insere(&lista, &tam, 14)
	fmt.Println(lista)

	// insere item em lista cheia
	fmt.Println("\nInsere item em lista cheia")
	Insere(&lista, &tam, 19)
	Insere(&lista, &tam, 20)
	Insere(&lista, &tam, 21)
	Insere(&lista, &tam, 22)
	fmt.Println(lista)

	// remove item existente
	fmt.Println("\nRemove item existente")
	Remove(&lista, &tam, 14)
	fmt.Println(lista)

	// remove item inexistente
	fmt.Println("\nRemove item inexistente")
	Remove(&lista, &tam, 14)
	fmt.Println(lista)

	// remove item em lista vazia
	fmt.Println("\nRemove item em lista vazia")
	Remove(&lista, &tam, 17)
	Remove(&lista, &tam, 2)
	Remove(&lista, &tam, 5)
	Remove(&lista, &tam, 7)
	Remove(&lista, &tam, 9)
	Remove(&lista, &tam, 12)
	Remove(&lista, &tam, 19)
	Remove(&lista, &tam, 20)
	Remove(&lista, &tam, 21)
	Remove(&lista, &tam, 22)
	fmt.Println(lista)
}

// Insere insere um valor em uma lista
func Insere(lista *[MaxTam]int, tamanho_atual *int, valor int) {
	if *tamanho_atual == MaxTam {
		fmt.Println("Overflow")
	} else if BuscaLinear1(*lista, *tamanho_atual, valor) != -1 {
		fmt.Println("Elemento já está na lista")
	} else {
		lista[*tamanho_atual] = valor
		*tamanho_atual++
	}
}

// Remove remove um valor em uma lista
func Remove(lista *[MaxTam]int, tamanho_atual *int, valor int) {
	if *tamanho_atual == 0 {
		fmt.Println("Underflow")
	} else {
		pos := BuscaLinear1(*lista, *tamanho_atual, valor)
		if pos == -1 {
			fmt.Println("Elemento não está na lista")
		} else {
			for i := pos; i < *tamanho_atual-1; i++ {
				lista[i] = lista[i+1]
			}
			lista[*tamanho_atual-1] = 0
			*tamanho_atual--
		}
	}
}

// Remove2 remove um valor em uma lista de forma mais legível. Uma boa prática é retornar os erros primeiro e depois o sucesso
func Remove2(lista *[MaxTam]int, tamanho_atual *int, valor int) string {
	if *tamanho_atual == 0 {
		return "Underflow"
	}

	pos := BuscaLinear1(*lista, *tamanho_atual, valor)

	if pos == -1 {
		return "Elemento não está na lista"
	}

	for i := pos; i < *tamanho_atual-1; i++ {
		lista[i] = lista[i+1]
	}
	lista[*tamanho_atual-1] = 0
	*tamanho_atual--
	return "Elemento removido"
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

// buscaBinaria busca um valor em uma lista
func buscaBinaria1(lista [MaxTam]int, n, valor int) int {
	inicio := 0
	fim := n - 1
	for inicio <= fim {
		meio := int((inicio + fim) / 2)
		if lista[meio] == valor {
			return meio
		}
		if lista[meio] < valor {
			inicio = meio + 1
		} else {
			fim = meio - 1
		}
	}
	return -1
}