package main

import "fmt"

const M = 4

func main() {
	var fila [M]int
	in, fim := -1, -1

	insere(&fila, &in, &fim, 4)
	fmt.Println(fila)
	insere(&fila, &in, &fim, 2)
	fmt.Println(fila)
	insere(&fila, &in, &fim, 3)
	fmt.Println(fila)
	remove(&fila, &in, &fim)
	fmt.Println(fila)
	insere(&fila, &in, &fim, 1)
	fmt.Println(fila)
	insere(&fila, &in, &fim, 5)
	fmt.Println(fila)
	insere(&fila, &in, &fim, 7) // Overflow
	fmt.Println(fila)

	remove(&fila, &in, &fim)
	fmt.Println(fila)
	remove(&fila, &in, &fim)
	fmt.Println(fila)
	remove(&fila, &in, &fim)
	fmt.Println(fila)
	remove(&fila, &in, &fim)
	fmt.Println(fila)
	remove(&fila, &in, &fim) // Underflow
	fmt.Println(fila)
}

func insere(f *[M]int, in *int, fim *int, novoValor int) {
	if *in == (*fim+1)%M {
		fmt.Println("Overflow")
		return
	}

	*fim++
	if *fim == M {
		*fim = 0
	}
	f[*fim] = novoValor

	if *in == -1 {
		*in = 0
	}
}

func remove(f *[M]int, in *int, fim *int) {
	if *in == *fim && *fim == -1 {
		fmt.Println("Underflow")
		return
	}

	f[*in] = 0
	if *in == *fim {
		*in, *fim = -1, -1
	} else {
		*in++
		if *in == M {
			*in = 0
		}
	}
}
