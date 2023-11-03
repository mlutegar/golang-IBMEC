/*
	Michel Lutegar D'Orsi Pereira
	14/08/2023
	ANOTADO
*/

package main

import "fmt"

func main() {
	// Array é uma coleção de dados do mesmo tipo, e possui tamanho fixo.
	// Declação explicita
	var filmes [5]string

	filmes[0] = "O Senhor dos Anéis"
	filmes[1] = "O poderoso chefão"
	filmes[2] = "Barbie"
	fmt.Println(filmes)

	// Declaração curta
	numeros := [4]int{0, 2, 4, 6}
	fmt.Println(numeros)

	// Slices são estruturas mais flexiveis
	var novosNumeros []int
	novosNumeros = numeros[1:] // vai até o final
	fmt.Println(novosNumeros)

	novosNumeros = numeros[1:3] // não inclui o elemento de indice 3
	fmt.Println(novosNumeros)

	novosNumeros = numeros[1:4] // ele para antes do índice 4
	fmt.Println(novosNumeros)

	// printf: é um print que permite formatar o texto dentro.
	fmt.Printf("%p", &numeros) // printando o enders	fmt.Println("")

	// Estrutura de repetição para arrays e para slices
	fmt.Println("--------------------")

	for i := 0; i < len(numeros); i++ {
		fmt.Println(numeros[i])
	}

	for indice, num := range numeros { // retorna sempre dois valores para cada iteração
		fmt.Println("Indice: ", indice, "- valor: ", num)
	}

	for _, num := range numeros { // não salva o indice
		fmt.Println("Valor: ", num)
	}

	fmt.Println("-------------------")
	fmt.Println(numeros)
	modificarArray(numeros) // não altera o array numeros!
	fmt.Println(numeros)

	fmt.Println(novosNumeros)
	modificarSlice(novosNumeros) // altera o slice original!
	fmt.Println(novosNumeros)

	fmt.Println(numeros)

	maisNumeros := criarSlice()
	fmt.Println(maisNumeros)
}

// Em Go, quando passamos um array como parâmetro criamos uma cópia deste array
func modificarArray(a [4]int) {
	a[0] = 999
}

func modificarSlice(s []int) {
	s[0] = 999
}

func criarSlice() []int {
	novoSlice := []int{10, 20, 30}
	return novoSlice
}
