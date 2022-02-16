package main

import "fmt"

func main() {
	var lado1, lado2, lado3, lado4, suma, producto int
	fmt.Println("Ingrese lado 1:")
	fmt.Scan(&lado1)
	fmt.Println("Ingrese lado 2:")
	fmt.Scan(&lado2)
	fmt.Println("Ingrese lado 3:")
	fmt.Scan(&lado3)
	fmt.Println("Ingrese lado 4:")
	fmt.Scan(&lado4)
	suma = lado1 + lado2
	producto = lado3 * lado4
	fmt.Println("la suma de los dos primeros es:", suma)
	fmt.Println("el producto de los dos ultimos es: ", producto)
}
