package main

import "fmt"

func main() {
	var num1, num2, suma, dif, producto, division float32

	fmt.Println("Ingrese primero numero:")
	fmt.Scan(&num1)
	fmt.Println("Ingrese segundo numero:")
	fmt.Scan(&num2)
	if num1 > num2 {
		suma = num1 + num2
		dif = num1 - num2
		fmt.Println("La suma de ambos numeros es: ", suma, " Y la diferencia: ", dif)

	} else {
		producto = num1 * num2
		division = num1 / num2
		fmt.Println("El producto entre ambos numeros es: ", producto, " Y la division: ", division)
	}
}
