package main

import (
	ctt "AC3/contato"
	"fmt"
)

func main() {
	var listaContato [5]ctt.Contato

	// Interface por linha de comando que permite a inclusão ou exclusão de contatos.
	var op int
	op = 100

	for op != 0 {
		fmt.Println("\nCONTATOS")
		fmt.Print("Escolha o que quer fazer: (0 para sair, 1 para adicionar, 2 para excluir último contato, 3 para editar email, 4 exibir contatos) ")

		fmt.Scanln(&op)

		switch op {
		case 1:
			var nome string
			var email string

			fmt.Print("Nome: ")
			fmt.Scanln(&nome)
			fmt.Print("Email: ")
			fmt.Scanln(&email)

			ctt.AdicionaContato(nome, email, &listaContato)
		case 2:
			ctt.ExcluiContato(&listaContato)
		case 3:
			for indice, contato := range listaContato {
				fmt.Println("Indice: ", indice, "| Contato: ", contato)
			}

			fmt.Print("\nQual é o indice da conta que deseja alterar? ")
			var i int
			fmt.Scanln(&i)

			if (listaContato[i] != ctt.Contato{}) {
				var emailNovo string

				fmt.Print("Qual é o novo email? ")
				fmt.Scanln(&emailNovo)

				ctt.EditaEmail(i, emailNovo, &listaContato)
			} else {
				fmt.Println("\nNão há contato cadastrado nesse indice!")
			}

		case 4:
			for indice, contato := range listaContato { // retorna sempre dois valores para cada iteração
				fmt.Println("Indice: ", indice, "| Contato: ", contato)
			}
		}
	}
}
