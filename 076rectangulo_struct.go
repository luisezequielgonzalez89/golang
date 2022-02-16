package main
import "fmt"

type Rectangulo struct {
	ladoMayor int
    ladoMenor int
}
func(lados *Rectangulo) cargaRectangulo(){
	fmt.Println("Ingrese lado Mayor: ")
	fmt.Scan(&lados.ladoMayor)
	fmt.Println("Ingrese lado menor: ")
	fmt.Scan(&lados.ladoMenor)
}
func(lados Rectangulo) perimetro() {
	var perimetro int
	var superficie int
	perimetro = (lados.ladoMayor + lados.ladoMenor) * 2
	superficie = lados.ladoMayor * lados.ladoMenor
	fmt.Println("El perimetro es: ", perimetro)
	fmt.Println("La superficie es: ", superficie)
}
func (lados Rectangulo) duplicar() {
	var duplicarmay int
	var duplicarmen int
	duplicarmay = lados.ladoMayor * 2
	duplicarmen = lados.ladoMenor * 2
	fmt.Println("Lado mayor duplicado: ", duplicarmay)
	fmt.Println("Lado menor duplicado: ", duplicarmen)
}
func main() {
	for {
		var opcion string
		var lados Rectangulo
		lados.cargaRectangulo()
		lados.perimetro()
		lados.duplicar()
		fmt.Println("Dese cargar otro rectangulo?[s/n]")
		fmt.Scan(&opcion)
		if opcion == "n"{
			break;
		}
	}
}