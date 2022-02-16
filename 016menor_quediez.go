package main

import "fmt"

func main() {
	var num1, num2, num3 int
	fmt.Println("Ingresar primer valor:")
	fmt.Scan(&num1)
	fmt.Println("Ingrese segundo valor:")
	fmt.Scan(&num2)
	fmt.Println("Ingresar tercer valor:")
	fmt.Scan(&num3)

	if num1 < 10 || num2 < 10 || num3 < 10 {
		fmt.Println("Alguno de los numeros es menor que diez: ", num1, num2, num3)
	} else {
		fmt.Println("No todos son menores que diez: ", num1, num2, num3)
	}
}
