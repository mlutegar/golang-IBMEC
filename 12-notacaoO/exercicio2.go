/*
ALGORITMO

n alunos

enquanto n > 0:
	recebe n
	para i=n; i > 0; i-- faca
		recebe um valor string de n carateres

*/

package main

import "fmt"

func main() {
	// declarando as variaveis
	var num_alunos, aux int

	fmt.Scan(&num_alunos)

	for num_alunos > 0 {
		alunos := make([]int, num_alunos)
		max := 0

		for i := num_alunos; i > 0; i-- {
			for j := 0; j < num_alunos; j++ {
				fmt.Scan(&aux)
				if aux == 1 {
					alunos[j]++
				}
			}
		}

		for _, v := range alunos {
			if max < v {
				max = v
			}
		}

		fmt.Println(max)

		fmt.Scan(&num_alunos)
	}
}
