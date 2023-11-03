package main

import "fmt"

func main() {
	var x *int
	var y int
	// var &z int não é possivel

	fmt.Println(x, y) // nill

	// x = 1 // não é possivel pois x é um ponteiro e armazena um endereço de memória
	y = 1

	fmt.Println(x, y) // nill 1

	x = &y // x recebe o endereço de memória de y

	fmt.Println(x, y)  // 0xc0000b4008 1
	fmt.Println(*x, y) // 1 1

}
