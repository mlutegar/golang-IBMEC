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

		for _, aluno := range alunos {
			if aluno > max {
				max = aluno
			}
		}

		fmt.Println(max)

		fmt.Scan(&num_alunos)
	}
}
