package main

import (
	"fmt"
)

func main() {
	var i, j string = "Hi", "hello"
	var k int = 20
	fmt.Print(i, j)
	fmt.Print("\n")
	fmt.Print(i, " ", j)
	fmt.Print("\n")
	fmt.Print(i, "\n", j, "\n")

	fmt.Println(i, j)

	fmt.Printf("i has the value: %v and type is: %T\n", i, i)
	fmt.Printf("j has the value: %v and type is: %T\n", j, j)
	fmt.Printf("k has the value: %v and type is: %T\n", k, k)
}
