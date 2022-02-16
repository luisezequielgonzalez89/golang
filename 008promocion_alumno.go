package main

import "fmt"

func main() {
	var nota1, nota2, nota3, promedio, suma float32
	fmt.Println("Ingrese la nota 1:")
	fmt.Scan(&nota1)
	fmt.Println("Ingrese la nota 2:")
	fmt.Scan(&nota2)
	fmt.Println("Ingrese la nota 3:")
	fmt.Scan(&nota3)
	suma = nota1 + nota2 + nota3
	promedio = suma / 3
	if promedio >= 7 {
		fmt.Println("Promocionado! ", promedio)
	}
}
