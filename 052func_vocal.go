package main

import (
	"fmt"
	"strings"
)

func Vocal(stg1 string) {
	res := strings.ToUpper(stg1)
	if res == "A" || res == "E" || res == "I" || res == "O" || res == "U" {
		fmt.Println("la letra ingresada es una vocal: ", res)
	} else {
		fmt.Println("La letra ingresada es: ", res, ", NO pertenece a las vocales")
	}
}

func main() {
	var letra string
	fmt.Println("Ingrese letra: ")
	fmt.Scan(&letra)
	Vocal(letra)
}
