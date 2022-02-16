package main

import "fmt"

func superficie(lado1, lado2 float32) float32 {
	var superficie float32
	superficie = lado1 * lado2
	return superficie
}
func main() {
	var lado1, lado2 float32
	fmt.Print("Rectangulo 1, ingrese lado 1: ")
	fmt.Scan(&lado1)
	fmt.Print("Ingrese lado 2: ")
	fmt.Scan(&lado2)
	sup1 := superficie(lado1, lado2)
	fmt.Print("Rectangulo 2, ingrese lado 1: ")
	fmt.Scan(&lado1)
	fmt.Print("Ingrese lado 2: ")
	fmt.Scan(&lado2)
	sup2 := superficie(lado1, lado2)
	if sup1 == sup2 {
		fmt.Print("Los dos rectangulos tienen la misma superficie | ", "Rectangulo 1 : ", sup1, " Rectangulo 2: ", sup2)
	} else {
		if sup1 > sup2 {
			fmt.Print("El rectangulo 1 tiene mayor superficie: ", sup1)
		} else {
			fmt.Print("El rectangulo 2 tiene mayor superficie: ", sup2)
		}
	}
}
