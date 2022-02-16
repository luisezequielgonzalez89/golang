package main

import "fmt"

type pais struct {
	nombre string
	cantidadhab int
}

func main(){
	var pais1, pais2, pais3 pais
	fmt.Println("Ingrese pais: ")
	fmt.Scan(&pais1.nombre)
	fmt.Println("Ingrese cantidad de habitantes: ")
	fmt.Scan(&pais1.cantidadhab)
	fmt.Println("Ingrese pais: ")
	fmt.Scan(&pais2.nombre)
	fmt.Println("Ingrese cantidad de habitantes: ")
	fmt.Scan(&pais2.cantidadhab)
	fmt.Println("Ingrese pais: ")
	fmt.Scan(&pais3.nombre)
	fmt.Println("Ingrese cantidad de habitantes: ")
	fmt.Scan(&pais3.cantidadhab)
	if pais1.cantidadhab > pais2.cantidadhab && pais1.cantidadhab > pais3.cantidadhab {
		fmt.Println("El pais con mayor habitantes es: ", pais1.nombre)
	}else {
		if pais2.cantidadhab > pais1.cantidadhab && pais2.cantidadhab > pais3.cantidadhab {
			fmt.Println("El pais con mayor habitantes es: ", pais2.nombre)
		}else {
			fmt.Println("el pais con mayor habitantes es: ", pais3.nombre)
		}
	}
}