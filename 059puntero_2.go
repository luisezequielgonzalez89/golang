package main

import "fmt"

func main() {
    var f int
    var pe *int
    pe = &f
    for *pe = 1; *pe <= 10; *pe = *pe + 1 {
        fmt.Println(f)  //se imprime ? ? ? ? ? ? ? ? ? ?
    }
}