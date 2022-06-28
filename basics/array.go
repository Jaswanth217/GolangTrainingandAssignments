package main

import (
	"fmt"
)

func main() {
	var jas1 = [4]int{24, 34, 44, 54}
	jas2 := [5]int{5, 6, 8, 11, 10}
	var jas01 = [...]int{9, 10, 11, 15, 16, 19, 12, 13} // var array_name = [length]datatype{values}
	jas02 := [...]int{2, 3, 4, 5, 6, 12, 14}            // var array_name = [...]datatype{values}

	var cars = [4]string{"Merc", "F150", "Volvo", "X3"}
	var j1 = [5]int{} //Array initialisation
	var j2 = [5]int{1, 2, 3}
	var j3 = [5]int{1, 2, 3, 4, 5}
	var j4 = [...]int{1, 2, 1, 3, 4, 5, 56, 67}

	fmt.Println(jas1)
	fmt.Println(jas2)
	fmt.Println(jas01)
	fmt.Println(jas02)
	fmt.Println(cars)
	fmt.Println(jas02[4]) //Access elements of an array
	jas1[3] = 21
	fmt.Println(jas1) //Change of element in an array

	fmt.Println(j1)
	fmt.Println(j2)
	fmt.Println(j3)
	fmt.Println(j4)

	var p1 = [7]int{4: 20, 6: 50} //Initialise specific element in an array
	fmt.Println(p1)

	fmt.Println(len(j4)) // Find the length of array
	fmt.Println(len(jas02))
	fmt.Println(len(p1))
	fmt.Println(cap(p1))
}
