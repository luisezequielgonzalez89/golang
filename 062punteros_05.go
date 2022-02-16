package main

import "fmt"

func mayorMenor(vec [5]int, pmayor *int, pmenor *int) {
    *pmayor = vec[0]
    *pmenor = vec[0]
    for f := 1; f < len(vec); f++ {
        if vec[f] > *pmayor {
            *pmayor=vec[f]
        } else {
            if vec[f] < *pmenor {
                *pmenor=vec[f]
            }
        }
    }
}


func main() {
    vec := [5]int{10, 22, 3, 44, 12}
    var ma, me int
    mayorMenor(vec, &ma, &me)
    fmt.Println("Mayor elemento del vector:", ma)
    fmt.Println("Menor elemento del vector:", me)
}