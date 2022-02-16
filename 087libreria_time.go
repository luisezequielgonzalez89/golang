package main

import (
    "fmt"
    "time"
)

func main() {
    navidad2016 := time.Date(2016, 12, 25, 0, 0, 0, 0, time.UTC)
    var ahora time.Time
    ahora = time.Now()
    fmt.Println(ahora)
    diferencia := ahora.Sub(navidad2016)
    fmt.Println("Cantidad de horas de diferencia:", diferencia.Hours())
}