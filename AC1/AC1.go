package main

import "fmt"

func main() {

	var x int
	fmt.Println("Insira um número para vê se ele é primo:")
	fmt.Scan(&x)
	e_primo(x)

	fmt.Println("-------------------")

	var y int
	fmt.Println("Insira a posição do número da sequência de fibbonaci que deseja:")
	fmt.Scan(&y)
	fmt.Println("O número é:", fibo(y))

	fmt.Println("-------------------")

	var z int
	fmt.Println("Insira o número do dia da semana que deseja:")
	fmt.Scan(&z)
	dia_da_semana(z)

	fmt.Println("-------------------")

	var k int
	fmt.Println("Insiro um ano para verificar se ele é bissexto:")
	fmt.Scan(&k)
	if e_bissexto(k) {
		fmt.Println("É um ano bissexto!")
	} else {
		fmt.Println("Não é bissexto.")
	}
}

// e_primo: verifica se o número é primo ou não, caso não for retornar todos os números que ele é divisivel
func e_primo(num int) {
	var divisores = 0

	fmt.Print("Os divisores são: ")
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			divisores++
			fmt.Print(i, ", ")
		}
	}

	fmt.Print("\n")

	if divisores == 2 {
		fmt.Println("É um número primo!")
	} else {
		fmt.Println("Não é um número primo.")
	}

}

// fibo: recebe um número e calcula o n-ésimo elemento da série de fibonacci
func fibo(num int) int {
	if num == 1 || num == 2 {
		return 1
	}

	return fibo(num-1) + fibo(num-2)
}

// dia_da_semana: recebe um número e fala qual dia da semana ele relaciona
func dia_da_semana(num int) {
	switch num {
	case 1:
		fmt.Println("Domingo")
	case 2:
		fmt.Println("Segunda")
	case 3:
		fmt.Println("Terça")
	case 4:
		fmt.Println("Quarta")
	case 5:
		fmt.Println("Quinta")
	case 6:
		fmt.Println("Sexta")
	case 7:
		fmt.Println("Sabado")
	default:
		fmt.Println("ERRO! O dia inserido não é válido.")
	}
}

// e_bissexto: recebe um ano e verifica se ele é bissexto ou não
func e_bissexto(ano int) bool {
	if ano%4 == 0 {
		if ano%100 == 0 && ano%400 != 0 {
			return false
		} else {
			return true
		}
	}
	return false
}
