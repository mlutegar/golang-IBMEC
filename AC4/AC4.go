/*
Procedimento hanoi (n, or, dest, trab){
	se n> 0 entao{
		hanoi(m-1,or,trab,dest)
		move disco n de or p/ dest
		hanoi(n-1, trab, dest, or)
	}
}

Move: peça 1: A -> B
Move peça n: X -> Y
...
Resolvido em t movimentos.

Não precisa de input.
*/

package main

import "fmt"

func main()  {
	qtdMovimento := hanoi(64, "A", "C", "B", 1)
	fmt.Println("\nForam necessários", qtdMovimento-1, "movimentos")
}

func hanoi(pecaAtual int, origem string, destino string, trabalho string, movimento int) int  {
	if pecaAtual>0 {
		movimento = hanoi(pecaAtual-1, origem, trabalho, destino, movimento)
		fmt.Println(movimento, "º Move o disco", pecaAtual, "de", origem,"para", destino)
		movimento++
		movimento = hanoi(pecaAtual-1, trabalho, destino, origem, movimento)
	}

	return movimento
}