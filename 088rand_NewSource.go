package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    var ale int
    for f := 0; f < 200; f++ {
        ale = r.Intn(101)
        fmt.Print(ale, "-")
    }	
}