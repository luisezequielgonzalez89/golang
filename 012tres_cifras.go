package main

import "fmt"

func main() {
	var num int
	fmt.Println("Ingrese el numero: ")
	fmt.Scan(&num)

	if num >= 10 && num <= 99 {
		fmt.Println("El numero ingresado tiene 2 digitos: ", num)
	} else {
		if num > 0 && num <= 9 {
			fmt.Println("El numero ingresado tiene una cifra: ", num)
		} else {
			if num >= 100 && num <= 999 {
				fmt.Println("El numero ingrsado tiene 3 cifras: ", num)
			} else {
				fmt.Println("Error! Valor fuera de rango!")
			}
		}

	}
}
