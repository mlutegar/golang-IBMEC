package main

import "fmt"

type No struct {
	Chave    int
	Esquerda *No
	Direita  *No
	Altura   int
}

type Arvore struct {
	Raiz *No
}

func imprimeArvore(a Arvore) {
	if a.Raiz != nil {
		imprimeSimetrico(a.Raiz)
	}
}

func imprimeSimetrico(n *No) {
	if n.Esquerda != nil {
		imprimeSimetrico(n.Esquerda)
	}
	fmt.Println(n.Chave)
	if n.Direita != nil {
		imprimeSimetrico(n.Direita)
	}
}

func main() {
	arv := Arvore{}

	fmt.Println("\nInserindo valores em uma árvore nula")
	fmt.Println(insereArvore(&arv, 3))
	imprimeArvore(arv)

	fmt.Println("\nInserindo valores já existentes em uma árvore existente")
	fmt.Println(insereArvore(&arv, 3))
	imprimeArvore(arv)

	fmt.Println("\nInserindo valores novos em uma árvore existente")
	fmt.Println(insereArvore(&arv, 7))
	fmt.Println(insereArvore(&arv, 1))
	fmt.Println(insereArvore(&arv, 6))
	fmt.Println(insereArvore(&arv, 5))
	fmt.Println(insereArvore(&arv, 9))
	fmt.Println(insereArvore(&arv, 2))
	imprimeArvore(arv)
}

// buscaArvore: Busca a partir do nó raiz, o valor passado como parâmetro e retorna: 0 se não tiver árvore, 1 se o valor for encontrado, 2 ou 3 se o valor não for encontrado e indicar o lado da árvore que deve ser inserido
func buscaArvore(opcao *int, no **No, valor int) {
	if *no == nil {
		*opcao = 0
		return
	}

	if (*no).Chave == valor {
		*opcao = 1
		return
	}

	if valor < (*no).Chave {
		if (*no).Esquerda == nil {
			*opcao = 2
			return
		}

		*no = (*no).Esquerda
		buscaArvore(opcao, no, valor)
	} else {
		if (*no).Direita == nil {
			*opcao = 3
			return
		}

		*no = (*no).Direita
		buscaArvore(opcao, no, valor)
	}
}

func insereArvore(arvore *Arvore, valor int) string {
	var opcao int
	no := arvore.Raiz

	buscaArvore(&opcao, &no, valor)

	if opcao == 1 {
		return "Valor já existe na árvore"
	}

	novoNo := &No{Chave: valor}
	novoNo.Chave = valor
	novoNo.Esquerda = nil
	novoNo.Direita = nil

	if opcao == 0 {
		arvore.Raiz = novoNo
		return "Valor inserido com sucesso"
	} else if opcao == 2 {
		no.Esquerda = novoNo
		return "Valor inserido com sucesso"
	} else {
		no.Direita = novoNo
		return "Valor inserido com sucesso"
	}
}
