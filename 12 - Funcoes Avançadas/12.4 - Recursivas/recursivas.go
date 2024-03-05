package main

import (
	"fmt"
)

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao - 2) + fibonacci(posicao - 1)
}

func main() {
	fmt.Println("Funcoes Recursivas!")

	posicao := uint(10)
	fmt.Println(posicao)
	fmt.Println(fibonacci(posicao))
}
