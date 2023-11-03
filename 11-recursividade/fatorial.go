/*
	Michel Lutegar D'Orsi Pereira
	30/08/2023
*/

package main

import "fmt"

func main() {
	fmt.Println(fatorialSemRecursividade(6), fatorialComRecursividade(6))
}

func fatorialSemRecursividade(num int) int {
	fat := 1
	for i := 1; i <= num; i++ {
		fat = fat * i
	}
	return fat
}

func fatorialComRecursividade(num int) int {
	if num == 0 || num == 1 {
		return 1
	}
	return num * fatorialComRecursividade(num-1)
}

// o que é grau de complexidade
