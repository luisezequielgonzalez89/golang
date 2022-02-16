package main

import "fmt"

func main() {
	var mat [3][4]int

	fmt.Println("A continuacion debe cargar la matriz:")

	for f := 0; f < 3; f++ {
		for c := 0; c < 4; c++ {
			fmt.Println("Inserte valor:")
			fmt.Scan(&mat[f][c])
		}
	}
	fmt.Println("Se completo la carga")
	fmt.Println(mat)
	may := mat[0][0]
	for f := 0; f < 3; f++ {
		for c := 0; c < 4; c++ {

			if mat[f][c] > may {
				may = mat[f][c]
			}
		}
	}
	fmt.Println("El valor mayor en la matriz es: ", may)
}
