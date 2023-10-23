package main

import "fmt"

// No é um nó de uma lista encadeada
type No struct {
	chave   int // chave do nó
	proxima *No // ponteiro para o próximo nó
}

// ListaEncadeada é uma lista encadeada
type ListaEncadeada struct {
	cabeca *No // ponteiro para o primeiro nó da lista
}

func main() {
	var l ListaEncadeada

	n1 := &No{chave: 2} // declara um nó e atribui a chave 2
	n2 := &No{chave: 4} // declara um nó e atribui a chave 4
	l.cabeca = n1       // insere o nó 1 no final da lista encadeada
	n1.proxima = n2     // insere o nó 2 no final da lista encadeada

	fmt.Println("\nInformações da lista encadeada:")
	fmt.Println(l)  // exibe a lista encadeada
	fmt.Println(n1) // exibe o nó 1
	fmt.Println(n2) // exibe o nó 2

	fmt.Println("\nExibindo a lista encadeada:")
	l.ExibirLista() // exibe a lista encadeada

	fmt.Println("\nBuscando um nó existente na lista encadeada:")
	busc := l.BuscarSimples(2) // busca o nó com chave 2 na lista encadeada
	fmt.Println(busc)          // exibe o nó que tem a chave 2

	fmt.Println("\nBuscando um nó inexistente na lista encadeada:")
	busc = l.BuscarSimples(1) // busca o nó com chave 1 na lista encadeada
	fmt.Println(busc)         // exibe o nó que tem a chave 1

	fmt.Println("\nInserindo um nó na lista encadeada:")
	l.InsercaoFinalLista(6) // insere o nó com chave 6 no final da lista encadeada
	l.ExibirLista()         // exibe a lista encadeada

	fmt.Println("\nBuscando um nó existente na lista encadeada ordenada:")
	buscOrdAnt, buscOrd := l.BuscaOrdenada(4)
	fmt.Println(buscOrdAnt, buscOrd)

	fmt.Println("\nBuscando o primeiro nó da lista encadeada ordenada:")
	buscOrdAnt, buscOrd = l.BuscaOrdenada(2) // busca o nó com chave 2 na lista encadeada ordenada
	fmt.Println(buscOrdAnt, buscOrd)         // exibe o nó que tem a chave 2

	fmt.Println("\nBuscando um nó inexistente na lista encadeada ordenada:")
	buscOrdAnt, buscOrd = l.BuscaOrdenada(5)
	fmt.Println(buscOrdAnt, buscOrd)

	fmt.Println("\nInserindo um nó na lista encadeada ordenada:")
	l.InserirOrdenadoLista(3) // insere o nó com chave 3 na lista encadeada ordenada
	l.ExibirLista()           // exibe a lista encadeada

	fmt.Println("\nRemovendo um nó na lista encadeada ordenada:")
	l.RemoverLista(3) // remove o nó com chave 3 na lista encadeada ordenada
	l.ExibirLista()   // exibe a lista encadeada
}

// ExibirLista exibe a lista encadeada
func (l *ListaEncadeada) ExibirLista() {
	no := l.cabeca // nó atual

	for no != nil { // enquanto o nó atual for diferente de nil
		fmt.Println(no.chave) // exibe a chave do nó atual
		no = no.proxima       // atualiza o nó atual para o próximo nó, pois o nó atual já foi exibido
	}
}

// BuscarSimples busca um nó na lista encadeada
func (l *ListaEncadeada) BuscarSimples(chave int) *No {
	no := l.cabeca // inicializa o nó com o primeiro nó da lista encadeada

	for no != nil { // enquanto o nó atual for diferente de nil; aqui não poderia ser "for no == nil && no.chave != chave"? Pois se o nó for nil, não tem como comparar a chave
		if no.chave == chave { // se a chave do nó atual for igual a chave buscada
			return no // retorna o nó atual
		}
		no = no.proxima // atualiza o nó atual para o próximo nó
	}
	return nil // retorna nil se não encontrar a chave
}

// InsercaoFinalLista insere um nó no final da lista encadeada
func (l *ListaEncadeada) InsercaoFinalLista(chave int) {
	novoNo := &No{chave: chave} // declara o nó que será inserido na lista encadeada, com a chave informada
	no := l.cabeca              // declara o primeiro nó da lista encadeada

	if no == nil { // se o nó atual for nil
		l.cabeca = novoNo // atualiza o nó atual para o novo nó
	} else {
		for no.proxima != nil { // enquanto o próximo nó for diferente de nil
			no = no.proxima // atualiza o nó atual para o próximo nó
		}
		no.proxima = novoNo // atualiza o próximo nó para o novo nó
	}
}

/* ALGORTIMO 16.2: InsercaoListaOrdenada

Programa BuscaOrdenada(lista, chave)
	chaveAnterior := nil
	no := lista.cabeca

	se no == nil então retorne chaveAnterior, no

	enquanto no.chave < chave faça
		chaveAnterior := no
		no := no.proxima
		se no == nil então retorne chaveAnterior, no

	se no.chave == chave então retorne chaveAnterior, no
	retorne chaveAnterior, nil
fimPrograma
*/

// BuscaOrdenada busca um nó na lista encadeada ordenada que retorna o nó anterior e o nó atual
func (l *ListaEncadeada) BuscaOrdenada(chave int) (*No, *No) {
	var chaveAnterior *No // declara um variável do tipo nó para armazenar o nó anterior
	no := l.cabeca        // declara o primeiro nó da lista encadeada

	if no == nil { // se o nó atual for nil
		return chaveAnterior, no // retorna o nó anterior e o nó atual
	}

	for no.chave < chave { // enquanto a chave do nó atual for menor que a chave buscada
		chaveAnterior = no // atualiza o nó anterior para o nó atual
		no = no.proxima    // atualiza o nó atual para o próximo nó
		if no == nil {     // se o nó atual for nil
			return chaveAnterior, no // retorna o nó anterior e o nó atual
		}
	}

	if no.chave == chave { // se a chave do nó atual for igual a chave buscada
		return chaveAnterior, no // retorna o nó anterior e o nó atual
	}

	return chaveAnterior, nil // retorna o nó anterior e nil
}

// InserirOrdenadoLista insere um nó na lista encadeada ordenada
func (l *ListaEncadeada) InserirOrdenadoLista(chave int) {
	ant, no := l.BuscaOrdenada(chave) // declara o nó anterior e o nó atual

	if no != nil { // se o nó atual for diferente de nil
		return // retorna nil
	}

	novoNo := &No{chave: chave} // declara o nó que será inserido na lista encadeada, com a chave informada

	if ant == nil { // se o nó anterior for nil
		novoNo.proxima = l.cabeca // atualiza o próximo nó para o primeiro nó da lista encadeada
		l.cabeca = novoNo         // atualiza o ponteiro do primeiro nó da lista encadeada para o novo nó
	} else { // se o nó anterior não for nil
		novoNo.proxima = ant.proxima // atualiza o próximo nó para o próximo nó do nó anterior
		ant.proxima = novoNo         // atualiza o próximo nó do nó anterior para o novo nó
	}
}

func (l *ListaEncadeada) RemoverLista(chave int) *No {
	ant, no := l.BuscaOrdenada(chave) // declara o nó anterior e o nó atual

	if no == nil { // verifica se o nó atual é nil
		return nil // se o nó atual for nil, ele retorna nil. Pois, não tem como remover um nó que não existe
	}

	if ant == nil { // verifica se o nó anterior é nil
		l.cabeca = no.proxima // se o nó anterior for nil, atualiza o primeiro nó da lista encadeada para o próximo nó do nó atual
	} else { // se o nó anterior não for nil
		ant.proxima = no.proxima // atualiza o próximo nó do nó anterior para o próximo nó do nó atual
	}

	return no // retorna o nó atual, pois ele foi removido da lista encadeada
}
