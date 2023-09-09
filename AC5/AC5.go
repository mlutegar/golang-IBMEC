// Escreva um algoritmo em pseudocódigo e implemente em Go os problemas abaixo:
// ▪ Dado um array de números inteiros positivos, considerado ordenado crescentemente, e um valor
// alvo, encontre um par de números no array cuja soma seja igual ao valor alvo. Se nenhum par for
// encontrado, retorne um valor (-1, -1) indicando que nenhum par foi encontrado. Resolva esse
// problema com um algoritmo cuja complexidade é O(n).

/*
ALGORITMO

array = [1, 2, 3, 4, 5]
alvo = 7

i = 0
j = len(array) - 1

enquanto i < j
	verifica se a soma de array[i] + array[j] é igual a alvo
		se for, retorna array[i] e array[j]
	se não for, verifica se a soma é maior que o alvo
		se for, decrementa j
		se não for, incrementa i

*/

package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5}
	alvo := 7

	fmt.Println(getPar(array, alvo))
}

func getPar(array []int, alvo int) (int, int) {
	var i, j int
	i = 0
	j = len(array) - 1

	for i < j{
		if (array[i] + array[j] == alvo) {
			return array[i], array[j]
		} else if (array[i] + array[j] > alvo) {
			j--
		} else {
			i++
		}
	}

	return -1, -1
}