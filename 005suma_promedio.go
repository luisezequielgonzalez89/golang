package main

import "fmt"

func main() {
	var valor1, valor2, valor3, valor4, promedio, suma int
	fmt.Println("Ingrese valor 1:")
	fmt.Scan(&valor1)
	fmt.Println("Ingrese valor 2:")
	fmt.Scan(&valor2)
	fmt.Println("Ingrese valor 3:")
	fmt.Scan(&valor3)
	fmt.Println("Ingrese valor 4:")
	fmt.Scan(&valor4)
	suma = valor1 + valor2 + valor3 + valor4
	promedio = suma / 4
	fmt.Println("la suma de los dos primeros es:", suma)
	fmt.Println("el promedio de los dos ultimos es: ", promedio)
}
