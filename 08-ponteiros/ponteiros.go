package main

import (
	"fmt"
	"strings"
)

type Pessoa struct{
	Nome string
	Email string
}

// Métodos começam com letra maiuscula e funções com letras minusculas
func (p *Pessoa) AlterarEmail(novoEmail string){
	p.Email = novoEmail
}

// Método 1, sem usar ponteiro. Ele duplica a lista para poder fazer a minpulação. Método não muito eficiente pois duplica memoria.
func adicionaPessoa1(p Pessoa, a [5]Pessoa) [5] Pessoa{
	for ind, pessoa := range a{
		if (pessoa == Pessoa{}){
			a[ind] = p
			break
		}
	}
	return a
}

// Método 2, usando ponteiro. É um método mais eficiente de adicionar elementos a uma array. Nesse caso ele não duplica a array, ele altera a array diretamente.
func adicionaPessoa2(p Pessoa, a *[5]Pessoa){
	for ind, pessoa := range a{
		if (pessoa == Pessoa{}){
			a[ind] = p
			break
		}
	}
}

func main(){
	var pessoas [5]Pessoa
	p1 := Pessoa{
		Nome: "aaa",
		Email: "bbb",
	}

	fmt.Println(pessoas)
	adicionaPessoa2(p1, &pessoas)
	fmt.Println(pessoas)

	x := 5
	y := x

	fmt.Println(x,y) // 5 5

	x = 6
	fmt.Println(x,y) // 6 5

	z := &x // z é um ponteiro, que aponta para x
	fmt.Println(z)
	fmt.Println(x, *z) // 6 6
	fmt.Println(&x) // referência

	var w *int
	fmt.Println(w) // nill

	// Ponteiros de ponteiros
	a := &z
	fmt.Println(a)
	fmt.Println(*a)
	fmt.Println(**a)

	mensagem := "Olá, Mundo"
	alteraMensagem(&mensagem)
	fmt.Println(mensagem)
}

func alteraMensagem(msg *string){
	// strings.ReplaceAll(stringOriginal string, textoProcurar string, textoSubstituir string)
	// "mundo" -> turma

	*msg = strings.ReplaceAll(*msg, "Mundo", "Turma")
}