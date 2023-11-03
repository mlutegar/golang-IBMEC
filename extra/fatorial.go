package main

import "fmt"

func main() {
	fmt.Println(fatorial(6))
}

func fatorial(num int) int {
	if num == 0 || num == 1 {
		return 1
	}
	return num * fatorial(num-1)
}
