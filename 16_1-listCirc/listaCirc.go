package main

import "fmt"

// Código para demonstrar uma lista circular
// Com uma função para a exibição dos nós em uma lista circular;
// Com uma função para a inserção de um nó no inicio da lista circular;
// Com uma função para a remoção do primeiro nó;

type No struct {
	chave   int
	proxima *No
}

type ListaCircular struct {
	cabeca *No
}

// Sequencia de testes de cada função
func main() {
	var l ListaCircular
	l.InserirNaPrimeiraPosicao(2)
	l.InserirNaPrimeiraPosicao(4)

	fmt.Println("\nExibindo a lista circular:")
	l.ExibirListaCircular()

	fmt.Println("\nInserindo um nó na lista circular:")
	l.InserirNaPrimeiraPosicao(6)
	l.ExibirListaCircular()

	fmt.Println("\nRemovendo um nó na lista circular:")
	l.RemoverPrimeiroNo()
	l.ExibirListaCircular()

	fmt.Println("\nRemovendo todos os nós:")
	l.RemoverPrimeiroNo()
	l.RemoverPrimeiroNo()
	l.ExibirListaCircular()
}

// ExibirListaCircular exibe a lista circular
func (l *ListaCircular) ExibirListaCircular() {
	no := l.cabeca
	for no != nil {
		fmt.Println(no)
		no = no.proxima
		if no == l.cabeca {
			break
		}
	}
}

// InserirNo insere um nó no inicio da lista circular
func (l *ListaCircular) InserirNaPrimeiraPosicao(ch int) {
	novoNo := &No{chave: ch}
	no := l.cabeca
	if no == nil {
		l.cabeca = novoNo
		novoNo.proxima = novoNo
	} else {
		for no.proxima != l.cabeca {
			no = no.proxima
		}
		no.proxima = novoNo
		novoNo.proxima = l.cabeca
		l.cabeca = novoNo
	}
}

// RemoverNo remove o primeiro nó da lista circular
func (l *ListaCircular) RemoverPrimeiroNo() {
	no := l.cabeca
	if no == nil {
		return
	}

	if no.proxima == l.cabeca {
		l.cabeca = nil
		return
	}

	for no.proxima != l.cabeca {
		no = no.proxima
	}
	no.proxima = l.cabeca.proxima
	l.cabeca = l.cabeca.proxima
}
