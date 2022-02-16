package main
import "fmt"

type Libro struct {
	codigo int
	titulo string
	autor string
}

func CargaLibros(libros *[4]Libro) {
	for f := 0; f < len(libros); f++ {
		fmt.Print("Ingrese codigo: ")
		fmt.Scan(&libros[f].codigo)
		fmt.Print("Ingrese titulo: ")
		fmt.Scan(&libros[f].titulo)
		fmt.Print("Ingrese autor: ")
		fmt.Scan(&libros[f].autor)
	}
}

func imprimir(libros [4]Libro) {
	fmt.Println(libros)
}

func Busqueda(libros [4]Libro) {
	var autor string
	var existe bool
	existe = false
	fmt.Println("Ingrese nombre de autor: ")
	fmt.Scan(&autor)
	for f := 0; f < len(libros); f++ {
		if autor == libros[f].autor {
			fmt.Println("Titulo: ", libros[f].titulo)
			existe = true
		}
	}
	if existe == false {
		fmt.Println("No hay libros de este autor")
	}
}

func main() {
	var libros [4]Libro
	CargaLibros(&libros)
	imprimir(libros)
	Busqueda(libros)

}