package main

import "fmt"

func main() {
	var precio, cantidad, total float32
	fmt.Println("Ingrese el precio del articulo:")
	fmt.Scan(&precio)
	fmt.Println("Cuantos lleva del mismos?:")
	fmt.Scan(&cantidad)
	total = precio * cantidad
	fmt.Println("Total a abonar:", total)

}
