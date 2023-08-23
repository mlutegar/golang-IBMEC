package contato

type Contato struct {
	Nome  string
	Email string
}

// AlterarEmail: função que altera o email do contato
func (c *Contato) AlterarEmail(emailNovo string) {
	c.Email = emailNovo
}