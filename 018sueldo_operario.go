package main

import "fmt"

func main() {
	var sueldo, antiguedad, ajuste, aumento float32
	fmt.Println("Ingrese sueldo:")
	fmt.Scan(&sueldo)
	fmt.Println("Ingrese antiguedad:")
	fmt.Scan(&antiguedad)
	fmt.Scan(aumento)

	if sueldo < 500 && antiguedad >= 10 {
		ajuste = (sueldo * 20) / 100
		aumento = sueldo + ajuste
		fmt.Print("El sueldo a pagar es: ", aumento)
	} else {
		if sueldo < 500 && antiguedad < 10 {
			ajuste = (sueldo * 5) / 100
			aumento = ajuste + sueldo
			fmt.Println("El sueldo a pagar es: ", aumento)
		} else {
			if sueldo >= 500 {
				fmt.Println("El sueldo es: ", sueldo)
			} else {
				fmt.Println("Gracias, vuelva pronto!")
			}
		}
	}
}
