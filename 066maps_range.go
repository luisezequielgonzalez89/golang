package main

import "fmt"

func main() {
    paises := map[string]int{
        "argentina" : 40000000,
        "españa"    : 46000000,
        "brasil"    : 190000000,
        "uruguay"   : 3400000,
    }
    for clave, valor := range paises {
        fmt.Println("Clave:", clave, " Valor:", valor)
    }
}