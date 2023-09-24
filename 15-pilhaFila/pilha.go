package main

import "fmt"

const M = 3

func main() {
	var pilha [M]int
	topo := 0

	insere(&pilha, &topo, 4)
	fmt.Println(pilha, topo)
	insere(&pilha, &topo, 5)
	fmt.Println(pilha, topo)
	insere(&pilha, &topo, 2)
	fmt.Println(pilha, topo)
	insere(&pilha, &topo, 1)
	fmt.Println(pilha, topo)

	remove(&pilha, &topo)
	fmt.Println(pilha, topo)
	remove(&pilha, &topo)
	fmt.Println(pilha, topo)
	remove(&pilha, &topo)
	fmt.Println(pilha, topo)
	remove(&pilha, &topo)
	fmt.Println(pilha, topo)
}

func insere(p *[M]int, topo *int, novoValor int) {
	if *topo == M {
		fmt.Println("Overflow")
		return
	}
	p[*topo] = novoValor
	*topo++
}

func remove(p *[M]int, topo *int) {
	if *topo == 0 {
		fmt.Println("Underflow")
		return
	}
	*topo--
	valorRemovido := p[*topo]
	p[*topo] = 0
	fmt.Println("Valor removido:", valorRemovido)
}