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
		fmt.Println("O que deseja fazer? \n 0: Sair \n 1: Adicionar \n 2: Excluir último contato \n 3: Editar algum email \n 4: Exibir contatos")

		fmt.Scanln(&op)

		switch op {
		case 1:
			fmt.Println("\nADIÇÃO DE CONTATO")
			var nome string
			var email string

			fmt.Print("Nome: ")
			fmt.Scanln(&nome)
			fmt.Print("Email: ")
			fmt.Scanln(&email)

			ctt.AdicionaContato(nome, email, &listaContato)
		case 2:
			fmt.Println("\nCONTATO EXCLUIDO")
			ctt.ExcluiContato(&listaContato)
		case 3:
			fmt.Println("\nEDIÇÃO DE EMAIL")
			for indice, contato := range listaContato {
				if (listaContato[indice] != ctt.Contato{}) {
					fmt.Println("Indice: ", indice, "| nome: ", contato.Nome, "| email: ", contato.Email)
				} else {
					fmt.Println("-")
				}
			}

			fmt.Print("Qual é o indice da conta que deseja alterar? ")
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
			fmt.Println("\nEXIBIÇÃO DE CONTATO")
			for indice, contato := range listaContato {
				if (listaContato[indice] != ctt.Contato{}) {
					fmt.Println("Indice: ", indice, "| nome: ", contato.Nome, "| email: ", contato.Email)
				} else {
					fmt.Println("-")
				}
			}
		}
	}
}
