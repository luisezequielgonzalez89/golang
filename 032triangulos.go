package main

import "fmt"

func main() {
	var lado1, lado2, lado3, cantidad, equilatero, isosceles, escaleno int

	fmt.Println("Ingrese la cantidad de triangulos que va a analizar: ")
	fmt.Scan(&cantidad)
	for count := 0; count <= cantidad; count++ {
		fmt.Println("Ingrese lado 1:")
		fmt.Scan(&lado1)
		fmt.Println("Ingrese lado 2:")
		fmt.Scan(&lado2)
		fmt.Println("Ingrese lado 3:")
		fmt.Scan(&lado3)
		if lado1 == lado2 && lado2 == lado3 {
			fmt.Println("El triangulo ingresado es equilatero")
			equilatero = equilatero + 1
		} else {
			if lado1 == lado2 || lado1 == lado3 || lado2 == lado3 {
				fmt.Println("El triangulo es isosceles")
				isosceles = isosceles + 1
			} else {
				if lado1 != lado2 && lado1 != lado3 {
					fmt.Println("El triangulo es escaleno")
					escaleno = escaleno + 1
				}
			}
		}
	}
	fmt.Println("Cantidad de triangulos equilateros: ", equilatero)
	fmt.Println("Cantidad de triangulos escalenos: ", escaleno)
	fmt.Println("Cantidad de triangulos isosceles: ", isosceles)
	if equilatero < escaleno && equilatero < isosceles {
		fmt.Println("Hay menor cantidad de equilatero: ", equilatero)
	} else {
		if escaleno < isosceles {
			fmt.Println("Hay menos cantidad de escaleno: ", escaleno)
		} else {
			fmt.Println("Hay menor cantidad de isosceles: ", isosceles)
		}
	}

}
