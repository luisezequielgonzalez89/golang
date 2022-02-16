package main

import "fmt"

func main() {
	var vector1 [10]int
	fmt.Println("Ingrese elementos del vector: ")

	for f := 0; f < len(vector1); f++ {
		fmt.Println("Ingrese valor: ")
		fmt.Scan(&vector1[f])
	}
	orden := 1
	for f := 0; f < len(vector1)-1; f++ {
		if vector1[f+1] < vector1[f] {
			orden = 0
		}
	}
	if orden == 1 {
		fmt.Println("Esta ordenado de menos a mayor")
	} else {
		fmt.Println("No esta ordenado de menor a mayor")
	}
}
