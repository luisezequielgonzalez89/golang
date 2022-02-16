package main

import (
    "fmt"
    "strings"
)

func main() {
    var palabra1, palabra2 string
    fmt.Print("Ingrese primer palabra:")
    fmt.Scan(&palabra1)
    fmt.Print("Ingrese segunda palabra:")
    fmt.Scan(&palabra2)
    pos := strings.Index(palabra1, palabra2)
    if pos == -1 {
        fmt.Println("No esta contenida la palabra", palabra2, "dentro de", palabra1)
    } else {
        fmt.Println("La palabra", palabra2, "esta contenida en", palabra1, "a partir de la posición", pos)
    }
}