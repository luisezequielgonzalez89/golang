package main

import "fmt"

func main() {
	var lista1, lista2, acumulado1, acumulado2 int
	x := 1
	for x <= 15 {
		fmt.Println("Ingrese valor de lista1:")
		fmt.Scan(&lista1)
		x = x + 1
		acumulado1 = acumulado1 + lista1
	}
	x = 1
	for x <= 15 {
		fmt.Println("Ingrese valor de lista2:")
		fmt.Scan(&lista2)
		x = x + 1
		acumulado2 = acumulado2 + lista2
	}
	if acumulado1 > acumulado2 {
		fmt.Println("Lista 1 mayor! ", acumulado1)
	} else {
		if acumulado1 < acumulado2 {
			fmt.Print("Lista 2 mayor! ", acumulado2)
		} else {
			fmt.Println("Listas iguales!")
		}
	}
}
