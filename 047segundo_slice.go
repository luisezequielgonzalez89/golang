package main

import "fmt"

func main() {
	var cantidad, prom, suma, mayor int
	fmt.Print("Indique la cantidad de enteros a cargar: ")
	fmt.Scan(&cantidad)

	slice := make([]int, cantidad)

	for f := 0; f < len(slice); f++ {
		fmt.Println("Ingrese valor entero:")
		fmt.Scan(&slice[f])
		suma = suma + slice[f]
	}
	prom = suma / cantidad
	fmt.Println("Los valores ingresados son: ", slice)
	for f := 0; f < len(slice); f++ {
		if slice[f] > prom {
			mayor = mayor + 1
		}
	}
	fmt.Println("Hay ", mayor, " valor/es mayor/es que el promedio!")
}
