package main

import "fmt"

func main() {
	terminos := 11
	sumando := 0
	x := 1

	for x <= 25 {
		x = x + 1
		sumando = terminos + sumando
		fmt.Println("Vamo sumando!! ", sumando)

	}
}
