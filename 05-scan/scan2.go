package main

import (
	"bufio" // um módulo que controla dados de entrada e saida
	"fmt"
	"os" //
)

func main() {
	var x int
	var y float64
	var texto string
	fmt.Println("Informe um número: ")
	fmt.Scanln(&x)
	fmt.Println(x)

	fmt.Println("Informe um float: ")
	fmt.Scanln(&y) // o Scanln ele ler a linha incluindo os \n, ou seja até a quebra de linha
	fmt.Println(y)

	fmt.Println("Informe uma palavra em texto: ")
	fmt.Scanln(&texto)
	fmt.Println(texto)

	leitor := bufio.NewReader(os.Stdin) // os.Stidin é o local onde o bufio vai ler, no caso Stdin é o terminal de controle do windows
	fmt.Println("Informe uma cadeia de texto: ")
	msg, _ := leitor.ReadString('\n') // ler até o bit '\n', ou seja quebra de linha

	// 'a' é diferente de "a", o '' se refere a bit e "" se refere a caracter

	fmt.Println(msg)
}
