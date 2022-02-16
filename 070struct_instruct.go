package main
import "fmt"

type Punto struct{
	x int
	y int
}

type Triangulo struct {
	punto1 Punto
	punto2 Punto
	punto3 Punto
}

func Carga(parametro *Triangulo) {
	fmt.Println("Ingrese x primer punto: ")
	fmt.Scan(&parametro.punto1.x)
	fmt.Println("Ingrese y para primer punto: ")
	fmt.Scan(&parametro.punto1.y)
	fmt.Println("Ingrese x segundo punto: ")
	fmt.Scan(&parametro.punto2.x)
	fmt.Println("Ingrese y para segundo punto: ")
	fmt.Scan(&parametro.punto2.y)
	fmt.Println("Ingrese x tercer punto: ")
	fmt.Scan(&parametro.punto3.x)
	fmt.Println("Ingrese y para tercer punto: ")
	fmt.Scan(&parametro.punto3.y)
}
func imprimir(parametro Triangulo) {
	fmt.Println("Los vertices del triangulos estan ubicados en: ")
	fmt.Println("Triangulo 1: ", parametro.punto1.x, parametro.punto1.y)
	fmt.Println("Triangulo 2: ", parametro.punto2.x, parametro.punto2.y)
	fmt.Println("Triangulo 3: ", parametro.punto3.x, parametro.punto3.y)
}
func main(){
	var triangulo1 Triangulo
	Carga(&triangulo1)
	imprimir(triangulo1)
}