package main

import "fmt"

// Código para demonstrar uma lista duplamente encadeada com os seus nós ordenados
// Com uma função para a exibição dos nós em uma lista duplamente encadeada;
// Com uma função para a busca de um nó;
// Com uma função para a inserção de um nó;
// Com uma função para a remoção de um nó;

type No struct {
	anterior *No
	chave    int
	proxima  *No
}

type ListaDuplamenteEncadeado struct {
	cabeca *No
}

// Sequencia de testes de cada função
func main() {
	var l ListaDuplamenteEncadeado
	l.InserirNo(2)
	l.InserirNo(4)

	fmt.Println("\nExibindo a lista duplamente encadeada:")
	l.ExibirListaDuplamenteEncadeada()

	fmt.Println("\nBuscando um nó na lista duplamente encadeada:")
	buscAnt, buscNo := l.BuscarNo(4)
	fmt.Println(buscAnt, buscNo)

	fmt.Println("\nBuscando um nó não existente na lista duplamente encadeada:")
	buscAnt, buscNo = l.BuscarNo(3)
	fmt.Println(buscAnt, buscNo)

	fmt.Println("\nInserindo um nó na lista duplamente encadeada:")
	l.InserirNo(3)
	l.InserirNo(10)
	l.InserirNo(1)
	l.ExibirListaDuplamenteEncadeada()

	fmt.Println("\nRemovendo um nó na lista duplamente encadeada:")
	l.RemoverNo(2)
	l.ExibirListaDuplamenteEncadeada()

	fmt.Println("\nRemovendo um nó não existente na lista duplamente encadeada:")
	l.RemoverNo(5)
	l.ExibirListaDuplamenteEncadeada()

	fmt.Println("\nRemovendo todos os nós:")
	l.RemoverNo(3)
	l.RemoverNo(4)
	l.RemoverNo(1)
	l.RemoverNo(10)
	l.ExibirListaDuplamenteEncadeada()
}

// ExibirListaDuplamenteEncadeada exibe a lista duplamente encadeada
func (l *ListaDuplamenteEncadeado) ExibirListaDuplamenteEncadeada() {
	no := l.cabeca
	for no != nil {
		fmt.Println(no)
		no = no.proxima
	}
}

// BuscarNo busca um nó na lista duplamente encadeada ordenada, que retorna o nó e o nó anterior
func (l *ListaDuplamenteEncadeado) BuscarNo(ch int) (*No, *No) {
	var ant *No
	no := l.cabeca

	if no == nil {
		return ant, no
	}

	for no.chave < ch {
		ant = no
		no = no.proxima
		if no == nil {
			return ant, no
		}
	}

	if no.chave == ch {
		return ant, no
	}

	return ant, nil
}

// InserirNo insere um nó na lista duplamente encadeada ordenada
func (l *ListaDuplamenteEncadeado) InserirNo(ch int) {
	ant, no := l.BuscarNo(ch)

	if no != nil {
		return
	}

	novoNo := &No{chave: ch}

	if l.cabeca == nil {
		l.cabeca = novoNo
		return
	}

	if ant == nil {
		novoNo.proxima = l.cabeca
		l.cabeca.anterior = novoNo
		l.cabeca = novoNo
		return
	}

	novoNo.proxima = ant.proxima
	ant.proxima = novoNo
	novoNo.anterior = ant
	if novoNo.proxima != nil {
		novoNo.proxima.anterior = novoNo
	}

}

// RemoverNo remove um nó na lista duplamente encadeada ordenada
func (l *ListaDuplamenteEncadeado) RemoverNo(ch int) {
	no := l.cabeca
	for no != nil {
		if no.chave == ch {
			if no.anterior == nil && no.proxima == nil {
				l.cabeca = nil
				break
			}

			if no.anterior == nil {
				l.cabeca = no.proxima
				no.proxima.anterior = nil
				break
			}

			if no.proxima == nil {
				no.anterior.proxima = nil
				break
			}

			no.anterior.proxima = no.proxima
			no.proxima.anterior = no.anterior

			break
		}
		no = no.proxima
	}
}
