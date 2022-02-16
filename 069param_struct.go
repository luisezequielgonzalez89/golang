package main
import "fmt"

type pais struct {
	nombre string
	cantidadhab int
}

func cargaDatos(parametro *pais){
	fmt.Println("Ingrese nombre de Pais: ")
	fmt.Scan(&parametro.nombre)
	fmt.Println("Ingrese cantidad de habitantes: ")
	fmt.Scan(&parametro.cantidadhab)
}

func mostrarPaises(parametro pais) {
	fmt.Println("Pais: ", parametro.nombre)
	fmt.Println("Cantidad de habitantes: ", parametro.cantidadhab)
}

func main(){
	var pais1, pais2, pais3 pais
	cargaDatos(&pais1)
	cargaDatos(&pais2)
	cargaDatos(&pais3)
	mostrarPaises(pais1)
	mostrarPaises(pais2)
	mostrarPaises(pais3)
}
