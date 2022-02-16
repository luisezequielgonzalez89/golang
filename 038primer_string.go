package main

import "fmt"

func main() {
	var name1, name2 string

	fmt.Println("Ingrese nombre 1:")
	fmt.Scan(&name1)
	fmt.Println("Ingrese nombre2:")
	fmt.Scan(&name2)

	if name1 > name2 {
		fmt.Println("El nombre1: ", name1, " es alfabeticamente mayor que el nombre 2: ", name2)
	} else {
		fmt.Println("El nombre 2: ", name2, " es alfabeticamente mayor que el nombre 1: ", name1)
	}
}
