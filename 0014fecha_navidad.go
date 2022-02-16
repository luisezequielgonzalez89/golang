package main

import "fmt"

func main() {
	var day, year int
	var mounth string
	fmt.Println("Por favor ingrese el dia con numeros: ")
	fmt.Scan(&day)
	fmt.Println("Por favor ingrese el mes con el nombre:")
	fmt.Scan(&mounth)
	fmt.Println("Por favor ingrese el ano con numeros:")
	fmt.Scan(&year)
	if day == 25 && mounth == "diciembre" {
		fmt.Println("Estamos en Navidad!! ", day, "/", mounth, "/", year)
	} else {
		fmt.Println("Hoy es ", day, " de ", mounth, " de ", year)
	}
}
