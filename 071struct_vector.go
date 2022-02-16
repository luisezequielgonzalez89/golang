package main
import "fmt"

type Producto struct {
	codigo int
	descripcion  string
	precio float64
}

func Carga(producto *[4]Producto) {
	for f := 0; f < len(producto); f++ {
		fmt.Println("Ingrese codigo: ")
		fmt.Scan(&producto[f].codigo)
		fmt.Println("Ingrese despcripcion: ")
		fmt.Scan(&producto[f].descripcion)
		fmt.Println("Ingrese precio: ")
		fmt.Scan(&producto[f].precio) 
	}
}

func mayor(producto [4]Producto) {
	pos := 0
	for f := 1; f < len(producto); f++ {
		if producto[pos].precio  < producto[f].precio {
			pos = f
		}
	}
	fmt.Println("Producto mas caro: ", producto[pos].descripcion)
}

func main(){
	var producto [4]Producto
	Carga(&producto)
	mayor(producto)
}