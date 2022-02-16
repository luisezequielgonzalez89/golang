package main
import "fmt"

type Empleado struct {
	nombre string
	sueldo float64
}
func(paramextra *Empleado) cargaEmpleado() {
	fmt.Println("Ingrese nombre de empleado: ")
	fmt.Scan(&paramextra.nombre)
	fmt.Println("Ingrese sueldo: ")
	fmt.Scan(&paramextra.sueldo)
}
func(paramextra Empleado) listarEmpleados() {
	fmt.Println("El empleado: ", paramextra.nombre, " tiene un sueldo de: ", paramextra.sueldo)
}
func(paramextra Empleado) impuestos() {
	if paramextra.sueldo > 3000 {
		fmt.Println("El empleado supera el sueldo promedio: ", paramextra.sueldo, " - Debe pagar impuestos.")
	}else {
		fmt.Println("El empleado no paga impuestos.")
	}
}
func main() {
	for{
		var sueldos Empleado
		var opcion string
		sueldos.cargaEmpleado()
		sueldos.listarEmpleados()
		sueldos.impuestos()
		fmt.Println("Desea cargar mas datos?[s/n]")
		fmt.Scan(&opcion)
		if opcion == "n" {
			break;
		}
	}
}