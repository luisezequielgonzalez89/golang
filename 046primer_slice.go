package main

import "fmt"

func main() {
	var cantidad int

	fmt.Print("Por favor ingrese la cantidad de productos a cargar:")
	fmt.Scan(&cantidad)
	producto := make([]string, cantidad)
	precio := make([]float64, cantidad)

	for f := 0; f < len(producto); f++ {
		fmt.Println("Ingrese nombre del producto:")
		fmt.Scan(&producto[f])
		fmt.Println("Ingrese precio de dicho producto:")
		fmt.Scan(&precio[f])
	}
	fmt.Println("Listado de productos ingresados: ")
	for f := 0; f < len(producto); f++ {
		fmt.Println(producto[f], " - ", precio[f])
	}
}
