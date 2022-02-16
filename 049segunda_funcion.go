package main

import (
	"fmt"
	"sort"
)

func mostrarMenor() {
	var cantidad int

	fmt.Println("Ingrese la cantidad de numeros que desea ingresar: ")
	fmt.Scan(&cantidad)

	menor := make([]int, cantidad)
	for f := 0; f < len(menor); f++ {
		fmt.Println("Ingrese valor: ")
		fmt.Scan(&menor[f])
	}
	sort.Ints(menor)
	fmt.Println("El menor de los valores ingresados es: ", menor[0])
	fmt.Println("Todos los valores ingresados: ", menor)
}
func main() {
	mostrarMenor()
}
