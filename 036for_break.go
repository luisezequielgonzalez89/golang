package main

import "fmt"

func main() {
	var valor, suma int

	for {
		fmt.Println("Igrese numero: ")
		fmt.Scan(&valor)
		if valor == 9999 {
			break
		}
		suma = valor + suma
	}
	if suma == 0 {
		fmt.Println("El valor ingresado es igual a 0: ", suma)
	} else {
		if suma < 0 {
			fmt.Println("El valor ingresado es menor a 0: ", suma)
		} else {
			fmt.Println(" el nummero es mayor a 0: ", suma)
		}
	}
}
