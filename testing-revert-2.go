package main

import "fmt"

func main() {
	var num1, num2, num3 int
	fmt.Println("Ingrese el num1:")
	fmt.Scan(&num1)
	fmt.Println("Ingrese el num2:")
	fmt.Scan(&num2)
	fmt.Println("Ingrese el num3:")
	fmt.Scan(&num3)
	if num1 > num2 && num1 > num3 {
		fmt.Println("El numero 1 es el mayor: ", num1)
	} else {
		if num2 > num3 {
			fmt.Println("El numero 2 es el mayor: ", num2)
		} else {
			fmt.Println("El numero 3 es el mayor: ", num3)
		}
	}