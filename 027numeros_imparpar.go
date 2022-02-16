package main

import "fmt"

func main() {
	var num, valores int
	x := 1
	fmt.Println("Ingrese cantidad de numeroa comprobar:")
	fmt.Scan(&num)

	for x <= num {
		x = x + 1
		fmt.Println("Ingresar valor:")
		fmt.Scan(&valores)
		if valores%2 == 0 {
			fmt.Println("Numero Par!")
		} else {
			fmt.Println("Numero Impar!")
		}
	}
}
