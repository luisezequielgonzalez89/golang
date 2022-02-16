package main

import "fmt"

func ingresoint() {
	var valor int
	fmt.Println("Ingrese valor: ")
	fmt.Scan(&valor)
	fmt.Println("El cuadrado del valor ingresado es: ", valor*valor)
}
func producto() {
	var valor1, valor2, producto int
	fmt.Println("Ingrese valor 1: ")
	fmt.Scan(&valor1)
	fmt.Println("Ingrese valor 2: ")
	fmt.Scan(&valor2)

	producto = valor1 * valor2
	fmt.Println("El prodcuto de entre ambos valores es: ", producto)

}
func main() {
	ingresoint()
	producto()
}
