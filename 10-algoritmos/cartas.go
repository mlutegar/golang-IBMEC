/*
	ALGORITMO - CARTAS

	n1, n2 := 0 // aqui poderiamos usar um int8, pois o intervalo de números é bem pequeno.
	classificação := "" // aqui a gente pode usar caracter (rune), em vez de usar uma string

	Para cada i entre 1 e 5, faça:
		Lê n2d
		se n1 = 0:
			entao n1 = n2
			continue
		se n2 > n1, então:
			se classificação = "" entao:
				classificação = C
			senao se classificação = D entao:
				classificação = N
				pare
		fazer análogo para se n1>n2
		n1 = n2

	escreve classificação
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	n1, n2 := 0, 0
	classificação := ""

	// Lê uma linha inserida pelo usuario
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	linha := scanner.Text()

	// Divide a linha em elementos, utilizando espaço como esperado
	// numeros é uma slice de strings
	numeros := strings.Split(linha, " ")

	for _, strNumero := range numeros {
		// O segundo retorno é uma mensagem de erros, caso dê algum problema
		// Estamos assumindo que não vai dar problema
		n2, _ = strconv.Atoi(strNumero)

		if n1 == 0{
			n1 = n2
			continue
		}

		if n2>n1{
			if classificação == ""{
				classificação = "C"
			} else if classificação == "D" {
				classificação = "N"
				break
			}
		}

		if n1>n2{
			if classificação == ""{
				classificação = "D"
			} else if classificação == "C" {
				classificação = "N"
				break
			}
		}

		n1 = n2

	}

	fmt.Println(classificação)
}