package main

import (
	"fmt"
	"strings"
)

func main(){
	cadena := "Esto es una prueba y es facil"
	slice1 := strings.Split(cadena, " ")
	for f := 0; f <len(slice1); f++ {
		fmt.Println(slice1[f])
	}
}
// package main

// import (
//     "fmt"
//     "strings"
// )

// func main() {
//     cadena := "Esto es una prueba y es fácil"
//     slice1 := strings.Split(cadena," ")
//     for f := 0; f < len(slice1); f++ {
//         fmt.Println(slice1[f])
//     }
// }