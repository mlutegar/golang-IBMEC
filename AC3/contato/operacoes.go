package contato

// AdicionaContato: recebe um contato e um array com até 5 elementos e inclui o contato no primeiro elemento vazio do array.
func AdicionaContato(nomeContato string, emailContato string, lista *[5]Contato){
	for contatoId, contatoValue := range lista {
		if (contatoValue == Contato{}) {
			lista[contatoId] = Contato{Nome: nomeContato, Email: emailContato}
			break
		}
	}
}

// ExcluiContato: recebe um array de 5 elementos e elimina o último contato válido.
func ExcluiContato(lista *[5]Contato){
	if (lista[4] != Contato{}) {
		lista[4] = Contato{}
	} else {
		for contatoId, contatoValue := range lista {
			if (contatoValue == Contato{}) {
				lista[contatoId-1] = Contato{}
				break
			}
		}
	}
}

// EditaEmail: recebe o índice do elemento no array e altera o e-mail do elemento indicado.
func EditaEmail(indice int, emailNovo string, lista *[5]Contato) {
	lista[indice].AlterarEmail(emailNovo)
}