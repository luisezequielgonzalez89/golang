package main

import "fmt"

func perimetro(valor1 float32) float32 {
	var perimetro float32
	perimetro = valor1 * 4
	return perimetro
}
func main() {
	var valor1 float32
	fmt.Print("Ingrese valor: ")
	fmt.Scan(&valor1)
	fmt.Print("El perimetro es: ", perimetro(valor1))
}
