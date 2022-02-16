package main

import "fmt"

func main() {
	var valor, num, negativos, positivos, pares, impar int

	fmt.Println("Por favor ingrese la cantidad de numero a verificar:")
	fmt.Scan(&valor)

	for count := 1; count <= valor; count++ {
		fmt.Println("Ingrese un numero:")
		fmt.Scan(&num)

		if num < 0 {
			negativos = negativos + 1
			if num%2 == 0 {
				pares = pares + 1
			} else {
				if num%15 == 0 {
					impar = impar + 1
				}
			}
		}
		if num > 0 {
			positivos = positivos + 1
			if num%2 == 0 {
				pares = pares + 1
			} else {
				if num%3 == 0 {
					impar = impar + 1
				}
			}
		}
	}
	fmt.Println("Cantidad de numero positivos: ", positivos)
	fmt.Println("Cantidad de numeros negativos: ", negativos)
	fmt.Println("Cantidad de numeros multiplos de 15: ", impar)
	fmt.Println("Cantidad de numeros pares: ", pares)
}
