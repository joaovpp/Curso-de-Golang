package main

import "fmt"

func main() {
	canal := make(chan string, 2)
	canal <- "Olá Mundo!"
	canal <- "Programando em Go!!"
	
	fmt.Println(<-canal)
	fmt.Println(<-canal)
}