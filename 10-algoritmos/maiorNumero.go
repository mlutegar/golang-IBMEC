/*
PROGRAMA MAIOR_NUMERO

	maior := 0
	enqunto veradeiro, faça:
		lê numero
		se numero = 0 entao pare
		se numero > maior então maior = numero
	escreve maior
*/

package main

import "fmt"

func main()  {
	var numero int
	maior := 0

	// caso fosse usar o menor número possivel teria que ser
	// var numero int32 // não pode ser int, pois o int depende da arquitetura do computador
	// maior := 2_147_483_648

	for{
		fmt.Scan(&numero) // vê qual é a diferença entre Scanln e Scan. O Scan só lê a linha e o Scanln lê a linha

		if numero == 0{
			break
		}

		if numero > maior{
			maior = numero
		}
	}

	fmt.Println(maior)
}