package main

import "fmt"

func main() {
	var num int
	fmt.Println("Ingrese el numero:")
	fmt.Scan(&num)
	if num > 9 {
		fmt.Println("El numero tiene 2 digitos")
	} else {
		fmt.Println("El numero tiene un solo digito")
	}
}
