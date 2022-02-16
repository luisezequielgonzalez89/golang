package main

import "fmt"

func main() {
	var altura, promedio, suma float32
	x := 1
	cantidad := 0
	fmt.Println("Ingrese la cantidad de alturas: ")
	fmt.Scan(&cantidad)

	for x <= cantidad {
		fmt.Println("Ingrese altura:")
		fmt.Scan(&altura)
		x = x + 1
		suma = altura + suma
		promedio = suma / float32(cantidad)
	}
	fmt.Println("La altura promedio es: ", promedio)
	fmt.Println("Se ingresaron, ", cantidad, "cantidad de alturas")
}
