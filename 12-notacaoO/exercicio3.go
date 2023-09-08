// Dado um array de números inteiros positivos, encontre o comprimento da maior subsequência crescente contígua. Uma subsequência crescente é uma sequência de elementos em que cada elemento subsequente é estritamente maior do que o anterior

package main

import (
	"fmt"
)

func main() {
	var array [12]int = [12]int{1, 2, 1, 1,3,67,123,45,2314,54,4, 5}

	tam := getMaiorSubsequênciaCresc(array)

	fmt.Println(tam)
}

func getMaiorSubsequênciaCresc(array [12]int) int {
	soma, somaAux := 1 , 1

	for i, _ := range array {
		if i == 0 {
			continue
		} else if array[i] >= array[i-1] {
			somaAux++
		} else {
			somaAux = 1
		}

		if somaAux > soma {
			soma = somaAux
		}
	}

	return soma
}