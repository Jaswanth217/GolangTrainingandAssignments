package main

import "fmt"

func main() {
	var primNum, primMin, primMax, primcount int

	fmt.Print("Enter the Minimum and Maximum Limit for Prime Number =")
	fmt.Scanln(&primMin, &primMax)

	fmt.Println("print number between ", primMin, "and", primMax, "are")
	for primNum = primMin; primNum <= primMax; primNum++ {
		primcount = 0
		for i := 2; i < primNum/2; i++ {
			if primNum%i == 0 {
				primcount++
				break
			}
		}
		if primcount == 0 && primNum != 1 {
			fmt.Print(primNum, "\t")
		}
	}
	fmt.Println()
}
