package main

import "fmt"

func valorEntero(valor1 int) {
	for f := 1; f <= valor1; f++ {
		if valor1 < 0 {
			fmt.Print("Valor negativo por favor ingrese un valor correcto:")
		} else {
			fmt.Println("Valor: ", f)
		}
	}
}

func main() {
	var x int
	fmt.Println("Por favor ingrese valor: ")
	fmt.Scan(&x)
	valorEntero(x)
}
