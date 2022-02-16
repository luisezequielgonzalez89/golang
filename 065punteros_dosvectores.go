package main
import "fmt"

func cargaProductos(vec1 *[5]string, vec2 *[5]int) {
	for f := 0; f < len(vec1); f++ {
		fmt.Println("Ingrese el nombre: ")
		fmt.Scan(&vec1[f])
		fmt.Println("Ingrese el precio: ")
		fmt.Scan(&vec2[f]) 
	}
}

func consultaPorNombre(vec1 [5]string, vec2 [5]int) {
	var match string
    existe := 0
	fmt.Println("Ingrese nombre del producto: ")
	fmt.Scan(&match)
	for f := 0; f < len(vec1); f++ {
		if match == vec1[f] {
			fmt.Println("El producto tiene un precio de: ", vec2[f])
			existe = 1
		}
	}
	if existe == 0 {
		fmt.Println("No existe un producto con dicho nombre almacenado")
	}
}

func main(){
	var vec1  [5]string
    var vec2  [5]int
	cargaProductos(&vec1, &vec2)
	consultaPorNombre(vec1, vec2)
}