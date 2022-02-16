package main

import (
	"fmt"
)

func main() {
	var h int = 19
	
	var p1 *int
	var p2 = new(int)
	p3 := &h
	fmt.Println("Variable v: ", h)
	// %T nos permite imprimir el tipo de dato de la variable
	fmt.Print("p1: %T \n", h)
	fmt.Printf("p2: %T \n", p2)
	fmt.Printf("p3: %T \n", p3)	
}