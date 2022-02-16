package main

import "fmt"

func main() {
    var z1, z2 float64
    var pf *float64
    pf = &z1
    *pf = 10.2
    pf = &z2
    *pf=20.90
    fmt.Println(z1, z2) // Se imprime ? ?
}