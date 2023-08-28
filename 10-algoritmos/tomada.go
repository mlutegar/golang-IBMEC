// Desafio tirado do br.spoj.com

package main

import "fmt"

func main(){
	var t1, t2, t3, t4 int

	fmt.Scanln(&t1, &t2, &t3, &t4)
	totalTomadas := t1+t2+t3+t4

	fmt.Println(totalTomadas - 3)
}