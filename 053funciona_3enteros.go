package main

import "fmt"

func tresEnteros(v1, v2, v3 int) int {
	var prom, suma int

	suma = v1 + v2 + v3
	prom = suma / 3

	return prom
}
func main() {
	var valor1, valor2, valor3 int
	fmt.Print("Ingrese valor 1: ")
	fmt.Scan(&valor1)
	fmt.Print("Ingrese valor 2: ")
	fmt.Scan(&valor2)
	fmt.Print("Ingrese valor 3: ")
	fmt.Scan(&valor3)
	fmt.Print("El promedio de los 3 numeros ingresados es: ", tresEnteros(valor1, valor2, valor3))
}
