package main

import "fmt"

func main() {
    var s1 string = "uno"
    var s2 string ="dos"
    var ps *string
    ps = &s1
    fmt.Println(s1) //se imprime: ?
    *ps = "tres"
    fmt.Println(s1) //se imprime: ?
    s1 = "cuatro"
    fmt.Println(*ps) //se imprime: ?
    ps = &s2
    fmt.Println(*ps) //se imprime: ?
}
