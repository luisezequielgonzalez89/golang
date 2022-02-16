package main

import "fmt"

func modificar(vec *[5]int) {
    for f := 0; f < len(vec); f++ {
        vec[f] = 10
    }
    fmt.Println("Contenido del parámetro de la funcón")    
    fmt.Println(vec) // 10 10 10 10 10
}

func main() {
    vec := [5]int{5,5,5,5,5}
    fmt.Println("Contenido del vector definido en la main en forma inicial") // 5 5 5 5 5 
    fmt.Println(vec)
    modificar(&vec)
    fmt.Println("Contenido del vector definido en la main luego de llamar a la función")  // 10 10 10 10 10
    fmt.Println(vec)
}