package main

import (
	"fmt"
)

func main() {
	var manana, tarde, noche, prom1, prom2, prom3 float32
	var suma1, suma2, suma3 float32

	for count := 1; count <= 5; count++ {
		/*manana = rand.Intn(35)*/
		fmt.Println("Ingrese edad turno manana:")
		fmt.Scan(&manana)
		fmt.Println("Edad turno mañana: ", manana)
		suma1 = manana + suma1

	}
	for count1 := 1; count1 <= 6; count1++ {
		fmt.Println("Ingrese edad turno tarde:")
		fmt.Scan(&tarde)
		fmt.Println("Edad turno tarde: ", tarde)
		suma2 = tarde + suma2
	}
	for count2 := 1; count2 <= 11; count2++ {
		fmt.Println("Ingrese edad turno noche:")
		fmt.Scan(&noche)
		fmt.Println("Edad turno noche: ", noche)
		suma3 = noche + suma3
	}
	prom1 = suma1 / 5
	prom2 = suma2 / 6
	prom3 = suma3 / 11
	if prom1 > prom2 && prom1 > prom3 {
		fmt.Println("El turno mañana tiene edades mayores promedio: ", manana)
	}
	if prom2 > prom1 && prom2 > prom3 {
		fmt.Println("El turno tarde tiene edades mayores  promedio: ", tarde)
	}
	if prom3 > prom2 && prom3 > prom1 {
		fmt.Println("El turno noche tiene edades mayores promedio: ", noche)
	}

	fmt.Println("Los promedios son: ", "mañana: ", manana, " tarde: ", tarde, " noche: ", noche)
}
