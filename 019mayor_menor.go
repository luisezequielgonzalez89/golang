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

	if num1 < num2 && num1 < num3 {
		fmt.Println("El menor es el: ", num1)
	} else {
		if num2 < num3 {
			fmt.Println("El menor es el: ", num2)
		} else {
			fmt.Println("El menor es el: ", num3)
		}
	}
	if num1 > num2 && num1 > num3 {
		fmt.Println("El numero mayor es: ", num1)
	} else {
		if num2 > num3 {
			fmt.Println("El numero mayor es: ", num2)
		} else {
			fmt.Println("El numero mayor es: ", num3)
		}
	}
}
