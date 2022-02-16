package main

import "fmt"

func main() {
	var cuenta, saldo int
	for {
		fmt.Println("Ingrese numero de cuenta:")
		fmt.Scan(&cuenta)
		fmt.Println("Ingrese saldo:")
		if cuenta <= 0 {
			fmt.Println("Carga finalizada")
			break
		}
		fmt.Scan(&saldo)
		if saldo > 0 {
			fmt.Println("Estado de la cuenta: \n")
			fmt.Println("Acreedor, Cuenta: ", cuenta, " Saldo: ", saldo)
		} else {
			if saldo < 0 {
				fmt.Println("Deudor, Cuenta: ", cuenta, " Saldo: ", saldo)
			} else {
				if saldo == 0 {
					fmt.Println("Nulo, Cuenta: ", cuenta, " Saldo: ", saldo)
				}
			}
		}
	}
}
