package main

import "fmt"

func main() {
	var x, y int
	fmt.Println("Ingrese coordenadas x:")
	fmt.Scan(&x)
	fmt.Println("Ingrese coordenadas y:")
	fmt.Scan(&y)
	if x > 0 && y < 0 {
		fmt.Println("Nos encontramos en el 1er cuadrante!")
	} else {
		if x < 0 && y > 0 {
			fmt.Println("Nos encontramos en el 2do cuadrante!")
		} else {
			fmt.Println("No reconozco este lugar! Estamos perdidos!")
		}
	}
}
