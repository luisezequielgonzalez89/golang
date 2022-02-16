package main

import "fmt"

func main() {
	var num, suma int

	for count := 0; count <= 10; count++ {
		fmt.Println("Favor de ingresar numero: ")
		fmt.Scan(&num)
		if count > 5 {
			suma = num + suma
		}
	}
	fmt.Println("La suma de los ultimos 5 es igual a: ", suma)
}
