package main

import(
	"fmt"
	"strings"
)

func main() {
	var cadena string
	var search string
	fmt.Println("Ingrese palabra: ")
	fmt.Scan(&cadena)
	fmt.Println("Ingrese cadena a buscar: ")
	fmt.Scan(&search)
	count := strings.Count(cadena, search)
	fmt.Println(count)
}