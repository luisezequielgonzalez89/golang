package main

import "fmt"

func digitos(vec [10]int) (int, int, int) {
	uno := 0
	dos := 0
	mas := 0

	for f := 0; f < 10; f++ {
		if vec[f] < 10 {
			uno++
		} else {
			if vec[f] >= 10 && vec[f] <= 99 {
				dos++
			} else {
				mas++
			}
		}
	}
	return uno, dos, mas
}

func main() {
	vector := [10]int{10, 399, 444, 2, 4, 5, 29, 3, 34, 45}
	v1, v2, v3 := digitos(vector)
	fmt.Println("Cantidad de elementos con 1 dígito:", v1)
	fmt.Println("Cantidad de elementos con 2 dígitos:", v2)
	fmt.Println("Cantidad de elementos con más de 2 dígitos:", v3)
}
