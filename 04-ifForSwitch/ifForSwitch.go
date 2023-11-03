/*
	Michel Lutegar D'Orsi Pereira
	09/08/2023
	ANOTADO
*/

package main

import "fmt"

func main() {
	var x = 10

	// ESTRUTURAS DE DECISÃO
	// if else e else if
	if x > 18 {
		fmt.Println("Caso 1")
	} else if x < 0 {
		fmt.Println("Caso 2")
	} else {
		fmt.Println("Caso 3")
	}

	// switch
	var dia = "segunda"
	switch dia {
	case "segunda", "terça", "quarta", "quinta", "sexta":
		fmt.Println("Dia de semana")
	case "sabado", "domingo":
		fmt.Println("Final de semana")
	default:
		fmt.Println("Erro, dia invalido")
	}

	// ESTRUTURAS DE REPETIÇÃO
	// For

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("-----------")

	x = 5
	// "while", não há o while originalmente
	for x > 0 {
		fmt.Println(x)
		x--
	}

	fmt.Println("-----------")

	// Uso do break e do continue
	for x < 10 {
		x++

		if x == 3 {
			continue
		}

		fmt.Println(x)

		if x == 5 {
			break
		}
	}
}
