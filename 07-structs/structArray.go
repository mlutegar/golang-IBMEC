/*
	Michel Lutegar D'Orsi Pereira
	14/08/2023
*/

package main

import "fmt"

type Pessoa struct {
	Nome string
}

func main() {
	var p = Pessoa{
		Nome: "Michel",
	}

	var pessoas [3]Pessoa

	pessoas[0] = p

	for _, pessoa := range pessoas {
		if (pessoa != Pessoa{}) {
			fmt.Println(pessoa)
		}
	}
}
