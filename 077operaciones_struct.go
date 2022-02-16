package main
import "fmt"

type Operacion struct {
	valor1 int
	valor2 int
}

func(valor *Operacion) cargaValores() {
	fmt.Println("Cargue valor 1: ")
	fmt.Scan(&valor.valor1)
	fmt.Println("Cargue valor 2: ")
	fmt.Scan(&valor.valor2)
}
func(valor Operacion) suma() {
	var operar int
	operar = valor.valor1 + valor.valor2
	fmt.Println("La operar de los valores es: ", operar)
}
func(valor Operacion) resta() {
	var operar int
	operar = valor.valor1 - valor.valor2
	fmt.Println("La resta de los valores es: ", operar)
}
func(valor Operacion) division() {
	var operar int
	operar = valor.valor1 / valor.valor2
	fmt.Println("La division de los valores es: ", operar)
}
func(valor Operacion) producto() {
	var operar int
	operar = valor.valor1 * valor.valor2
	fmt.Println("El producto de los valores es: ", operar)
}

func main() {
	for {
		var	 operaciones Operacion
		var opcion string
		operaciones.cargaValores()
		operaciones.suma()
		operaciones.resta()
		operaciones.division()
		operaciones.producto()
		fmt.Println("Desea realizar mas operaciones?:[s/n]")
		fmt.Scan(&opcion)
		if opcion == "n"{
			break;
		}
	}
}
