package main

import "fmt"

func main() {

	var valor, tabla int

	fmt.Println("Ingrese un valor para enseñarle la tabla de multiplicar del mismo:")
	fmt.Scan(&valor)

	for count := 0; count <= 12; count++ {
		if valor <= 10 && valor > 0 {
			tabla = valor + tabla
			fmt.Println("Tabla del ", valor, ":", tabla)
		} else {
			fmt.Println("Por favor ingrese un valorr comprendido entre 1 y 10: ")
			fmt.Scan(&valor)
		}
	}
}
