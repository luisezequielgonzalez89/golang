package main

import "fmt"

func main() {
	var num1, num2, num3, operacion int
	fmt.Println("Ingresar primer valor:")
	fmt.Scan(&num1)
	fmt.Println("Ingrese segundo valor:")
	fmt.Scan(&num2)
	fmt.Println("Ingresar tercer valor:")
	fmt.Scan(&num3)

	if num1 == num2 && num2 == num3 {
		operacion = (num1 + num2) * num3
		fmt.Println("El resultado es: ", operacion)
	} else {
		fmt.Println("Los numeros ingreador son: ", num1, num2, num3)
	}
}
