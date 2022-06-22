package main

import "fmt"

const (
	y = iota
	x
	r
	t
	q
	A
	n
	b
	z = 217
	c = 45
	d = "HI!"
	e
	f
	gh = "34 say"
	fg
	df
)

func main() {
	fmt.Println(y, x, r, t, q, A, n, b)
	fmt.Println(z, d)
	fmt.Println(gh)
	fmt.Println(y, A)
	fmt.Println(A, n, b)
	fmt.Println(c, d, e)
	fmt.Println(f, gh)
	fmt.Println(fg, df, b)
	fmt.Println(z, y, x, r, t, q)

}
