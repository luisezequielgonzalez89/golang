package main
import "fmt"
 
type Ciudadano struct {
	colegio string
    nromesa int
}

func carga(colegios map[int]Ciudadano) {
	var dni int
	var opcion string
	var mesa Ciudadano
	for {
		fmt.Println("Ingrese DNI: ")
		fmt.Scan(&dni)
		fmt.Println("Ingrese colegio: ")
		fmt.Scan(&mesa.colegio)
		fmt.Println("Ingrese nromesa: ")
		fmt.Scan(&mesa.nromesa)

		colegios [dni] = mesa
		
		fmt.Print("Desea Cargar nuevo votante?[s/n]? ")
		
		fmt.Scan(&opcion)
		if opcion == "n" {
			break;
		}
	}
}

func listar(colegios map[int]Ciudadano){
	for dni, mesa := range colegios{
		fmt.Println("Dni: ", dni)
		fmt.Println("colegio: ", mesa.colegio)
		fmt.Println("Numero de mesa: ", mesa.nromesa)
	}
}

func consulta(colegios map[int]Ciudadano) {
	var dni int
	fmt.Println("Ingrese DNI de votante: ")
	fmt.Scan(&dni)
	mesa, existe :=  colegios[dni]
	if existe {
		fmt.Println("Colegio: ", mesa.colegio)
		fmt.Println("Mesa: ", mesa.nromesa)
	}else {
		fmt.Println("No existe")
	}
}
func main() {
	Colegios := make(map[int]Ciudadano)
	carga(Colegios)
	listar(Colegios)
	consulta(Colegios)
}
