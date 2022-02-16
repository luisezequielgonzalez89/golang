package main

import "fmt"

func main() {
	var notas int
	x := 1
	mayor := 0
	menor := 0

	for x <= 10 {
		fmt.Println("Ingrese nota: ")
		fmt.Scan(&notas)
		x = x + 1
		if notas >= 7 {
			mayor = mayor + 1
		} else {
			menor = menor + 1
		}
	}
	fmt.Println("La cantidad de notas mayores que 7 son: ", mayor)
	fmt.Println("La cantidad de notas menor que 7 son: ", menor)
}
