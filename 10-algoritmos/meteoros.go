/*
PROGRAMA METEOROS
	contFazenda := 0
	mensagem := ""
	Enquanto VERDADEIRO, faça:
		Lê x1, y1, x2, y2
		Se x1=y1=x2=y2=0 então Pare
		contFazenda++
		mensagem += "Teste" + contFazenda
		Lê numMeteoritos
		meteoritosNaFazenda := 0
		Para cada meteorito entre 1 e numMeteoritos, faça:
			Lê x, y1
			Se x1<=x<=x2 e y2<=y<=y1, então:
				meteoritosNaFazenda++
		mensagem += meteoritosNaFazenda
	Escreve mensagem
*/

package main

import (
	"fmt"
	"strconv"
)

func main(){
	// declaração de variaveis
	var x, y, x1, y1, x2, y2, numMeteoritos, meteoritosNaFazenda int
	contFazenda := 0
	mensagem := ""

	for{
		fmt.Scanln(&x1, &y1, &x2, &y2)
		// Condições de parada é importante colocar sempre no inicio do for. É uma boa prática.
		if (x1 == 0 && x2 == 0 && y1 == 0 && y2 == 0){
			break
		}

		contFazenda++
		// strconv: é uma biblioteca usada para converter tipos inteiros com string, ou vircer versa
		mensagem += "Teste " + strconv.Itoa(contFazenda) + "\n"

		fmt.Scanln(&numMeteoritos)
		meteoritosNaFazenda = 0

		for meteorito := 1; meteorito <= numMeteoritos; meteorito++ {
			fmt.Scanln(&x, &y)
			if (x1 <= x && x <= x2 && y2 <= y && y <= y1) {
				meteoritosNaFazenda++
			}
		}

		mensagem += strconv.Itoa(meteoritosNaFazenda) + "\n\n"
	}

	fmt.Print(mensagem)
}