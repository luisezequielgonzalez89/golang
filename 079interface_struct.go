package main
import "fmt"

type Cuadrado struct {
	lado int
}
type Rectangulo struct {
	ladoMay int
	ladoMen int
}
type Figura interface {
	Perimetro()
	Superficie()
}
func(param Cuadrado)Perimetro() {
	fmt.Println("Perimetro de Cuadrado: ",param.lado * 4)
}
func(param Cuadrado)Superficie() {
	fmt.Println("Superficie de Cuadrado: ",param.lado * 2)
}
func(param Rectangulo)Perimetro() {
	fmt.Println("Perimetro de Rectangulo: ",(param.ladoMay * 2) + (param.ladoMen * 2))
}
func(param Rectangulo)Superficie() {
	fmt.Println("Superficie de Rectangulo: ",param.ladoMay * param.ladoMen)
}

func RecorrerSup(vect [2]Figura) {
	for	f := 0; f < len(vect); f++ {
			vect[f].Superficie()
		}
}
func RecorrerPem(vect [2]Figura) {
	for	f := 0; f < len(vect); f++ {
			vect[f].Perimetro()
		}
}
func main() {
	var vect [2]Figura
	vect [0] = Cuadrado{lado: 5}
	vect [1] = Rectangulo{ladoMay:12, ladoMen:7}
	RecorrerSup(vect)
	RecorrerPem(vect)
}
