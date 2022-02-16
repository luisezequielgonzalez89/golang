package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Println("Ingresar primer valor:")
	fmt.Scan(&num1)
	fmt.Println("Ingrese segundo valor:")
	fmt.Scan(&num2)
	if num1 > num2 {
		suma := num1 + num2
		resta := num1 - num2
		fmt.Println("La suma de ambos numeros es: ", suma, " - Y su resta es: ", resta)
	} else {
		if num1 < num2 {
			producto := num1 * num2
			division := float32(num1) / float32(num2)
			fmt.Println("El producto entre ambos numeros es: ", producto, " - Su division: ", division)
		} else {
			fmt.Println("Por favor ingrese un valor correcto")
		}
	}
}
