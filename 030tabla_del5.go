package main

import "fmt"

func main() {
	for count := 5; count <= 50; count = count + 5 {
		fmt.Println("Tabla del 5: ", count)
	}
}
