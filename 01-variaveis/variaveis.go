package main

import "fmt"

var linguagem string
var soma bool

func main() {
	// declaração explicita
	var msg string // Não pode criar variaveis que não vamos utilizar.
	var n1, n2 int

	// declaração implicita
	var n3, n4 = 15, 25

	linguagem = "go"

	// declaração curta
	msg2 := "ola, mundo"

	msg = "oi"
	n1 = 10
	n2 = 20

	fmt.Println(msg, n1, n2, n3, n4, msg2)

	const Pi = 3.14
	fmt.Println(linguagem)

	f2()

	var numero float64
	numero = 4.7


	var n5 = 10.55 // float64
	fmt.Println(numero, n5)

	var n6 complex128
	n6 = 18 + 25.4i
	fmt.Println(n6)
}

func f2() {
	linguagem = "python"
	fmt.Println(linguagem)
}

/*
TIPOS DE DADOS

- Inteiros
int8 -128 a 127
uint8 (uncient int) 0 a 255
int16 -32.768 a 32.767
uint16 0 a 65.535
int32
uint23
int64
uint64
int (ele vai ser 32 ou 64)

byte (igual ao uint8)
rune (igual a int32 mas especificado para caracteres unicode) (é mais usado para web)

- Flutuantes

float32	precisão de 6 a 9 digitos
float64	precisõ de 15 a 17 digitos

*não existe o float sozinho, é preciso especificar
*se for uma declaração implicita (que não precisa tipificar a variavel, ele entende como float64)

- Complexo
complex64
complex128

- String
bool	true / false
nil		ausência de dados

*/
