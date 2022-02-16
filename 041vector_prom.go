package main

import "fmt"

func main() {
	var vector1 [5]float32
	var vector2 [5]float32
	var prom1, prom2, sum1, sum2 float32

	fmt.Println("Ingrese las notas del curso A: ")
	for f := 0; f < len(vector1); f++ {
		fmt.Println("Curso A: ")
		fmt.Scan(&vector1[f])
		sum1 = sum1 + vector1[f]
	}
	fmt.Println("Ingrese las notas del curso B: ")
	for f := 0; f < len(vector2); f++ {
		fmt.Println("Ingrese notas del curso B: ")
		fmt.Scan(&vector2[f])
		sum2 = sum2 + vector2[f]
	}
	prom1 = sum1 / 5
	prom2 = sum2 / 5

	if prom1 > prom2 {
		fmt.Println("El curso con mayor promedio general es el A: ", prom1)
	} else {
		if prom1 == prom2 {
			fmt.Println("Ambos cursos tienen el mismo promedio - curso A: ", prom1, " Curso B: ", prom2)
		} else {
			fmt.Println("El curso con mayor promedio es el B: ", prom2)
		}
	}

}
