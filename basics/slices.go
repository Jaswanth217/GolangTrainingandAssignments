package main

import (
	"fmt"
)

func main() {
	j1slice := []int{}
	fmt.Println(len(j1slice))
	fmt.Println(cap(j1slice))
	fmt.Println(j1slice)

	j2slice := []string{"Audi", "BMW", "Volvo", "Alpharomeo", "TESLA"}
	fmt.Println(len(j2slice))
	fmt.Println(cap(j2slice))
	fmt.Println(j2slice)
}
