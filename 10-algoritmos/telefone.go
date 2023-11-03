/*
	Michel Lutegar D'Orsi Pereira
	30/08/2023
*/

package main

import "fmt"

func main() {
	// A função `make()` aloca memória para um slice de strings chamado `expressoes`.
	// O tamanho do slice é 0, o que significa que ele está inicialmente vazio.

	// Um slice é uma estrutura de dados que armazena uma sequência de elementos
	// de um tipo específico. Os slices são dinâmicos, o que significa que
	// seu tamanho pode ser alterado durante a execução do programa.

	// A notação []string é usada para declarar um slice de strings.
	var expressoes []string = make([]string, 0)

	for {
		var expressao string
		fmt.Scanln(&expressao)

		if expressao == "" {
			break
		}

		expressoes = append(expressoes, expressao)
	}

	for _, expressao := range expressoes {
		fmt.Println(encontreOTelefone(expressao))
	}
}

func encontreOTelefone(expr string) string {
	telefone := ""

	for _, carac := range expr {
		switch carac {
		case '1', '0', '-':
			telefone += string(carac)
		case 'A', 'B', 'C':
			telefone += "2"
		case 'D', 'E', 'F':
			telefone += "3"
		case 'G', 'H', 'I':
			telefone += "4"
		case 'J', 'K', 'L':
			telefone += "5"
		case 'M', 'N', 'O':
			telefone += "6"
		case 'P', 'Q', 'R', 'S':
			telefone += "7"
		case 'T', 'U', 'V':
			telefone += "8"
		case 'W', 'Y', 'X', 'Z':
			telefone += "9"
		}
	}

	return telefone
}
