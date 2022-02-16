package main

import "fmt"

func PositivosNegativos() (int, int) {
	var num int
	pos := 0
	neg := 0
	for {
		fmt.Println("Igrese numero: | (Para finalizar presiones 0)")
		fmt.Scan(&num)
		if num == 0 {
			break
		} else {
			if num < 0 {
				neg++
			} else {
				pos++
			}
		}
	}
	return neg, pos
}

func main() {
	neg, pos := PositivosNegativos()
	fmt.Println("La cantidad de valores positivos es: ", pos)
	fmt.Println("La cantidad de valores negativos es: ", neg)
}
