// Dado um array de números inteiros positivos e um valor alvo, encontre um par de números no array cuja soma seja igual ao valor alvo. Se nenhum par for encontrado, retorne um valor (-1, -1) indicando que nenhum par foi encontrado.

package main

import (
	"fmt"
)

func main() {
	numeroAlvo := 8
	var array [5]int = [5]int{1, 2, 3, 4, 5}

	num1, num2 := getPares(array, numeroAlvo)

	fmt.Println(num1, num2)
}

func getPares(array [5]int, valor int) (int, int){ // O(n^2)
	for _, i := range array {
		for _, j := range array {
			if (i+j == valor) {
				return i, j
			}
		}
	}

	return 0, 0
}