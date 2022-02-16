package main

import "fmt"

func main() {
	var preguntas, respuestas, porcentaje int
	fmt.Println("Ingrese cantidad de preguntas realizadas: ")
	fmt.Scan(&preguntas)
	fmt.Println("Cuantas fueron correctas?: ")
	fmt.Scan(&respuestas)
	porcentaje = respuestas * 100 / preguntas
	if porcentaje >= 90 {
		fmt.Println("Nivel Maximo alcanzado: ", porcentaje, "%")
	} else {
		if porcentaje >= 75 && porcentaje < 90 {
			fmt.Println("Nivel medio alcanzado: ", porcentaje, "%")
		} else {
			if porcentaje >= 50 && porcentaje < 75 {
				fmt.Println("Nivel regular alcanzado: ", porcentaje, "%")
			} else {
				if porcentaje < 50 {
					fmt.Println("Fuera de nivel, lo siento: ", porcentaje, "%")
				} else {
					fmt.Println("Por favor ingrese un valor numerico entero")
				}
			}
		}
	}
}
