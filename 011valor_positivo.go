package main

import "fmt"

func main() {
	var num int
	fmt.Println("Ingresar numero: ")
	fmt.Scan(&num)
	if num > 0 {
		fmt.Println("El numero ingresado es positivo", num)
	} else {
		if num == 0 {
			fmt.Println("El numero ingresado es nulo: ", num)
		} else {
			fmt.Println("El numero ingresado es negativo: ", num)
		}
	}
}
