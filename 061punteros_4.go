package main

import "fmt"

func intercambiar(pe1 *int, pe2 *int) {
     aux := *pe1
    *pe1 = *pe2
    *pe2 = aux
}

func main() {
    x1 := 10
    x2 := 20
    fmt.Println(x1, x2)
    intercambiar(&x1, &x2)
    fmt.Println(x1, x2)
}