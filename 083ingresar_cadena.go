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
	strings.Contains(cadena, search)
	if strings.Contains(cadena, search) {
		fmt.Println("La cadena ", search, " esta contenida en: ", cadena)
	}else {
		fmt.Println("La cadena ", search, " no esta contenida en ", cadena)
	}
}