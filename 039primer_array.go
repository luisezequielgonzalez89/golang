package main

import "fmt"

func main() {
	var arreglo [8]int
	var may int
	var total int
	var may1 int

	for f := 0; f < len(arreglo); f++ {
		fmt.Println("Ingrese elemento:")
		fmt.Scan(&arreglo[f])
		total = total + arreglo[f]

		if arreglo[f] > 36 && arreglo[f] <= 50 {
			may = may + arreglo[f]
		} else {
			if arreglo[f] > 50 {
				may1++
			}
		}
	}
	fmt.Println("El valo acumulado de todos los elementos es: ", total)
	fmt.Println("El valor acumulado de los valores mayores a 36: ", may)
	fmt.Println("La cantidad de valores mayores a 50 es: ", may1)
}
