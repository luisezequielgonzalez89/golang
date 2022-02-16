package main

import "fmt"

func main() {
	var x, y float32
	var puntos int
	fmt.Println("Por favor ingrese la cantidad de puntos en el plano a ingresar: ")
	fmt.Scan(&puntos)
	for count := 1; count <= puntos; count++ {
		fmt.Println("Ingrese coordenada x:")
		fmt.Scan(&x)
		fmt.Println("Ingrese coordenada y:")
		fmt.Scan(&y)
		if x > 0 && y > 0 {
			fmt.Println("Sus coordenadas se encuentra en el 1: ", "x= ", x, " y= ", y)
		} else {
			if x < 0 && y > 0 {
				fmt.Println("Sus coordenadas se encuentran en el cuadrante 2: ", "x=", x, " y= ", y)
			} else {
				if x < 0 && y < 0 {
					fmt.Println("Sus coordenadas se encuentran en el cuadrante 3: ", "x= ", x, " y= ", y)
				} else {
					if x > 0 && y < 0 {
						fmt.Println("Sus coordenadas se encuentran en el cuadrante 4: ", "x= ", x, " y= ", y)
					} else {
						if x == 0 && y == 0 {
							fmt.Println("Se encuentra en el punto de partida: ", "x= ", x, " y= ", y)
						}
					}
				}
			}
		}
	}
}
