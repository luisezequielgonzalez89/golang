package main

import "fmt"

func main() {
	var sueldos, importe, suma, suma1 float32
	mayor := 0
	menor := 0
	x := 1
	cantidad := 0
	fmt.Println("Ingrese la cantidad de empleados:")
	fmt.Scan(&cantidad)

	for x <= cantidad {
		x = x + 1
		fmt.Println("Ingresar sueldo:")
		fmt.Scan(&sueldos)
		if sueldos >= 100 && sueldos <= 300 {
			suma = suma + sueldos
			menor = menor + 1
		} else {
			if sueldos > 300 {
				suma1 = suma1 + sueldos
				mayor = mayor + 1
			}
		}
	}
	importe = suma + suma1
	fmt.Println("La cantidad de sueldos entre $100 y $300 son: ", menor)
	fmt.Println("La cantidad de sueldos mayores a $300 son: ", mayor)
	fmt.Println("La empresa invierte en sus empleados: ", importe)
}
