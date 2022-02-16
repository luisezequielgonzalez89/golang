package main

import(
	"fmt"
	"strings"
)

func main() {
	var cadena string
	var search string
	var  slice [2]string
	for f := 0; f < len(slice); f++{
		fmt.Println("Ingrese palabra: ")
		fmt.Scan(&cadena)
		fmt.Println("Ingrese valores a buscar: ")
		fmt.Scan(&search)
		first := strings.HasPrefix(cadena, search)
		if first == true {
			fmt.Println("El string buscado: ", search, " coincide con la palabra ingresada: ", cadena)
		}else {
			fmt.Println("El string buscado: ", search, " no coincide con la palabra ingresada: ", cadena)
		}
	}
}