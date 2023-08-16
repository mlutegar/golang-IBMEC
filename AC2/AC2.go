package main

import "fmt"

type Contato struct{
	nome string
	email string
}

// alterarEmail: função que altera o email do contato
func (c *Contato) alterarEmail(emailNovo string){
	c.email = emailNovo
}

func main()  {
	var listaContato [5]Contato

	// Interface por linha de comando que permite a inclusão ou exclusão de contatos.
	var op int
	op = 100

	for op != 0{
		fmt.Println("\n\nCONTATOS")
		fmt.Println("Escolha o que quer fazer: (0 para sair, 1 para adicionar, 2 para excluir último contato)")

		fmt.Scanln(&op)

		switch op {
		case 0:
			break
		case 1:
			var nome string
			var email string

			fmt.Print("Nome: ")
			fmt.Scanln(&nome)
			fmt.Print("Email: ")
			fmt.Scanln(&email)

			listaContato = adicionaContato(nome, email, listaContato)
			break
		case 2:
			listaContato = excluiContato(listaContato)
			break
		}

		fmt.Println(listaContato)
	}
}

// adicionaContato: recebe um contato e um array com até 5 elementos e inclui o contato no primeiro elemento vazio do array.
func adicionaContato(nomeContato string, emailContato string, lista [5]Contato) [5]Contato{

	for contatoId,contatoValue := range lista{
		if (contatoValue == Contato{}) {
			lista[contatoId] = Contato{nome: nomeContato, email: emailContato}
			break
		}
	}
	return lista

}

// excluiContato, que recebe um array de 5 elementos e elimina o último contato válido.
func excluiContato(lista [5]Contato) [5]Contato{
	for contatoId,contatoValue := range lista{
		if (contatoValue == Contato{}) {
			lista[contatoId-1] = Contato{}
			break
		}
	}
	return lista
}

