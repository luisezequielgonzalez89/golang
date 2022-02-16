package main

import "fmt"

func main() {
	var mat [2][5]int
	fmt.Println("Realice la carga de la matriz:")

	for f := 0; f < 2; f++ {
		for c := 0; c < 5; c++ {
			fmt.Println("Carga de matriz:")
			fmt.Scan(&mat[f][c])
		}
	}
	for f := 0; f < 2; f++ {
		for c := 0; c < 5; c++ {
			fmt.Println(mat)
		}
	}
}
