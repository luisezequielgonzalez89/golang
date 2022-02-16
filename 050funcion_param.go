package main

import "fmt"

func tresEnteros(valor1, valor2, valor3 int) {
	if valor1 > valor2 && valor1 > valor3 {
		fmt.Print("El valor mayor es: ", valor1)
	} else {
		if valor2 > valor1 && valor2 > valor3 {
			fmt.Print("El valor mayor es: ", valor2)
		}
		if valor3 > valor1 && valor3 > valor2 {
			fmt.Print("El valor mayor es: ", valor3)
		}
	}
}

func main() {
	var valor1, valor2, valor3 int
	fmt.Print("Ingrese valores: ")
	fmt.Print("Valor 1: ")
	fmt.Scan(&valor1)
	fmt.Print("Valor 2: ")
	fmt.Scan(&valor2)
	fmt.Print("Valor 2: ")
	fmt.Scan(&valor3)
	tresEnteros(valor1, valor2, valor3)
}
