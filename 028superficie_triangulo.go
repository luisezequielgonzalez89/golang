package main

import "fmt"

func main() {
	var base, altura, superficie float32
	var suma int
	for count := 0; count <= 5; count++ {
		fmt.Println("Ingrese la base del triangulo:")
		fmt.Scan(&base)
		fmt.Println("Ingrese la altura:")
		fmt.Scan(&altura)
		superficie = base * altura
		fmt.Println("La medida de la base es: ", base, " Su alturea: ", altura, " y superficie: ", superficie)
		if superficie > 12 {
			suma++
		}
	}
	fmt.Println("La cantida de triangulos con superficie mayor a 12 es: ", suma)
}

/*En este programa me di cuenta que si declaro una variabla adentro de un IF solo existe alli */
