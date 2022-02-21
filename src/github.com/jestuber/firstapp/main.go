package main

import (
	"fmt"
)

func main() {
	var n bool
	fmt.Printf("%v, %T\n", n, n)
	fmt.Println(n == false)

	m := 42
	fmt.Printf("%v, %T\n", m, m)
	m = 1e10

	var k uint16 = 42
	fmt.Printf("%v, %T\n", k, k)

	p := 3.14
	p = 13.7e72 //this is too large for float32, so go initializes p as a float64
	p = 2.1e14
	fmt.Printf("%v, %T\n", p, p)

	q := 3.14
	q = 2.1e14 //even though this could be a float32, go always initializes floats as float64 unless specified
	fmt.Printf("%v, %T\n", q, q)

	var i complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", real(i), real(i))
	fmt.Printf("%v, %T\n", imag(i), imag(i))

	var t complex128 = complex(5, 12)
	fmt.Printf("%v, %T\n", t, t)
}
