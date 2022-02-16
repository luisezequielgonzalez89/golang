package main
import {
	"fmt"
	"strings"
}
func cadenas() {
	cadena := "Esto es una prueba y es facil"
	slice1 := strings.Split(cadena, " ")
	for f := 0; f < len(slice1); f++ {
		count := 0
		if slice1[f] == "es"{
			count = count + 1
		}
	}
	fmt.Println("La palabra es se repite: ", count, " veces")
}