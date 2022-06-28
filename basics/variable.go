package main

import (
	"fmt"
)

var m int
var n int = 13
var o = 14

func main() {
	var j, a, s, u int = 12, 13, 14, 15
	var w, l = 16, "pj"
	t, h := 17, "pelaprolujaswanth"
	var (
		jas int
		wan int    = 18
		th  string = "GIS 217"
	)
	var firstname string = "Pelaprolu"
	var lastname = "Jaswanth"
	Age := 26
	Place := "India"
	var hi string = "This is just sample"
	var years int = 26
	var months float32 = 26.4
	var monthsInTotal int = 316                //Camel Case
	var months_total = "three hundred and six" //Snake Case
	var MonthInTotal int = 312                 //Pascal Case
	var torf bool

	m = 12
	fmt.Println(w, l)
	fmt.Println(t, h)
	fmt.Println(jas, wan, th)
	fmt.Println(m, n, o)
	fmt.Println(j, a, s, u)
	fmt.Println(firstname)
	fmt.Println(lastname)
	fmt.Println(Age)
	fmt.Println(Place)
	fmt.Println(firstname, lastname, Age)
	fmt.Println(hi)
	fmt.Println(years)
	fmt.Println(months)
	fmt.Println(monthsInTotal)
	fmt.Println(months_total)
	fmt.Println(MonthInTotal)
	fmt.Println(torf)
}
