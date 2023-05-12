package main
import (
	"fmt"
	"strconv"
)
func main() {
	// var dato string
	var slice [5]int
	var convert string
	for f := 0; f < len(slice); f++ {
		fmt.Println("Ingrese valor numerico: ")
		fmt.Scan(&slice[f])
		convert = strconv.Itoa(slice[f])
		fmt.Printf("Valor convertido a string: ", convert)
		fmt.Println("")
	}
}

/testing revert