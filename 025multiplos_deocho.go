package main

import "fmt"

func main() {
	terminos := 8
	sumando := 0

	for sumando < 500 {

		sumando = terminos + sumando
		fmt.Println("Vamo sumando!! ", sumando)

	}
}
