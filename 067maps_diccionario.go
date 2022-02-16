package main

import "fmt"

func carga(diccionario map[string]string){
	var ingles, castellano, opcion string
	for {
		fmt.Println("Ingrese palabra en ingles: ")
		fmt.Scan(&ingles)
		fmt.Println("Ingrese palabra traducida a castellano: ")
		fmt.Scan(&castellano)
		diccionario[ingles] = castellano
		fmt.Println("Desea continuar cargando palabras?[s/n]: ")
		fmt.Scan(&opcion)
		if opcion == "n"{
			break;
		}
	}
}
func listar(diccionario map[string]string){
	for ingles, castellano := range diccionario{
		fmt.Println("Ingles: ", ingles, " castellano: ", castellano)
	}
}
func buscar(diccionario map[string]string){
	var ingles string
	fmt.Println("Ingrese palabra a traducir: ")
	fmt.Scan(&ingles)
	_, existe := diccionario[ingles]
	if existe {
		fmt.Println("La palabra traducida significa: ", diccionario[ingles])
	}else {
		fmt.Println("La palabra ingresada no existe")
	}
}

func main(){
	diccionario := make(map[string]string)
	carga(diccionario)
	listar(diccionario)
	buscar(diccionario)

}