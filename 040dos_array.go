package main

import "fmt"

func main() {
	var vector1 [4]int
	var vector2 [4]int
	var vectorsum [4]int

	fmt.Println("Carga el primer vector: ")
	for f := 0; f < len(vector1); f++ {
		fmt.Println("Ingrese valor en vector 1: ")
		fmt.Scan(&vector1[f])
	}
	fmt.Println("Ingrese el segundo vector: ")
	for f := 0; f < len(vector2); f++ {
		fmt.Println("Ingrese valor en vector 2: ")
		fmt.Scan(&vector2[f])
	}
	for f := 0; f < len(vectorsum); f++ {
		vectorsum[f] = vector1[f] + vector2[f]
	}
	fmt.Println("Vector resultante: ", vectorsum)
}
