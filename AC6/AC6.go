package main

import (
	"fmt"
	"strconv"
)

const MaxValor = 5

func main() {
	var lista [MaxValor]int
	var n = 0

	fmt.Println(insereOrd(&lista, &n, 4))
	fmt.Println(lista)

	fmt.Println(insereOrd(&lista, &n, 12))
	fmt.Println(lista)

	fmt.Println(insereOrd(&lista, &n, 2))
	fmt.Println(lista)

	fmt.Println(insereOrd(&lista, &n, 6))
	fmt.Println(lista)

	fmt.Println(insereOrd(&lista, &n, 17))
	fmt.Println(lista)

	fmt.Println(insereOrd(&lista, &n, 1))
	fmt.Println(lista)
}

func insereOrd(lista *[MaxValor]int, tamanhoAtual *int, novoValor int) string{
	if MaxValor == *tamanhoAtual {
		fmt.Println("\nTentando inserir", novoValor)
		return "Overflow"
	}

	if buscaLinear(*lista, *tamanhoAtual, novoValor) == -1 {
		indice := *tamanhoAtual - 1

		for indice >= 0 && lista[indice] > novoValor {
			lista[indice+1] = lista[indice]
			indice--
		}

		fmt.Println("\nTentando inserir", novoValor)
		lista[indice+1] = novoValor
		*tamanhoAtual++
		return strconv.Itoa(novoValor) + " não encontrado, inserindo com sucesso"
	}

	fmt.Println("\nTentando inserir", novoValor)
	return "Elemento já existe"
}

func buscaLinear(lista [MaxValor]int, tamanhoAtual, valor int) int {
	indice := 0
	lista[tamanhoAtual] = valor

	for lista[indice] != valor {
		indice++
	}

	if indice < tamanhoAtual {
		return indice
	}

	return -1
}