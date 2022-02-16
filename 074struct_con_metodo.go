package main
import "fmt"

type Triangulo struct {
	lado1 int
	lado2 int
	lado3 int
}

func(paramextra *Triangulo) carga() {
	fmt.Println("Ingrese lado1: ")
	fmt.Scan(&paramextra.lado1)
	fmt.Println("Ingrese lado2: ")
	fmt.Scan(&paramextra.lado2)
	fmt.Println("Ingrese lado3: ")
	fmt.Scan(&paramextra.lado3)
}
func (paramextra Triangulo) ladoMayor() {
	if paramextra.lado1 > paramextra.lado2 && paramextra.lado1 > paramextra.lado3 {
		fmt.Println("El lado mayor es el Lado 1: ", paramextra.lado1)
	}else{
		if paramextra.lado2 > paramextra.lado1 && paramextra.lado2 > paramextra.lado1 {
			fmt.Println("El lado mayor es el Lado 2: ", paramextra.lado2)
		}else {
			fmt.Println("El lado mayor es el Lado 3: ", paramextra.lado3)
		}
	}
}
func (paramextra Triangulo) tipoTriangulo()  {
	if paramextra.lado1 == paramextra.lado2 && paramextra.lado1 == paramextra.lado3{
		fmt.Println("El triangulo SI es equilatero")
	}else{
		fmt.Println("El triangulo NO es equilatero")
	}
}

func main() {
var triangulo Triangulo
triangulo.carga()
triangulo.ladoMayor()
triangulo.tipoTriangulo()
}