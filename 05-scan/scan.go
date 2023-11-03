/*
	Michel Lutegar D'Orsi Pereira
	09/08/2023
	ANOTADO
*/

package main

import "fmt"

func main() {
	var x int
	var y float64
	var texto string
	fmt.Println("Informe um número: ")
	fmt.Scan(&x)
	fmt.Println(x)

	fmt.Println("Informe um float: ")
	fmt.Scan(&y)
	fmt.Println(y)

	fmt.Println("Informe um texto: ")
	fmt.Scan(&texto)
	fmt.Println(texto)
}
