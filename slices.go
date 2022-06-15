package main

import "fmt"

func main() {
	s := []int{2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Length of Slice: %d \n", len(s))
	fmt.Printf("Capacity of Slice: %d \n", cap(s))

	s = s[:0]
	fmt.Println(s)
	fmt.Printf("Length of Slice: %d \n", len(s))
	fmt.Printf("Capacity of Slice: %d \n", cap(s))

	s = s[1:4]
	fmt.Println(s)
	fmt.Printf("Length of Slice: %d \n", len(s))
	fmt.Printf("Capacity of Slice: %d \n", cap(s))

	s = s[2:]
	fmt.Println(s)
	fmt.Printf("Length of Slice: %d \n", len(s))
	fmt.Printf("Capacity of Slice: %d \n", cap(s))

	s = s[2:4]
	fmt.Println(s)
	fmt.Printf("Lenght of Slice: %d \n", len(s))
	fmt.Printf("Capacity of Slice: %d \n", cap(s))
}
