package main

import "fmt"

func main() {
	productinput()
}

func productinput() {
	var	productos = map[string]int{"Arroz": 10,"Arvejas": 5, "Leche": 15, "Lentejas": 4, "Pan": 2}
	fmt.Println(productos)
	for key, value := range productos {
	fmt.Println("Ingrese Nombre de producto: ")
	fmt.Scan(&key)
	fmt.Println("Ingres precio: ")
	fmt.Scan(&value)
	}
for key, value := range productos{
	fmt.Println(key, value)
}
}