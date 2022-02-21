package main

import (
	"fmt"
)

func main() {
	s := "this is a string"
	fmt.Printf("%v, %T\n", s, s)
	fmt.Printf("%v, %T\n", s[2], s[2])
	fmt.Printf("%v, %T\n", string(s[2]), s[2])

	s2 := "this is also a string"
	fmt.Printf("%v, %T\n", s+s2, s+s2)

	b := []byte(s) // byte slice
	fmt.Printf("%v, %T\n", b, b)

	r := 'a' //runes are type aliases for int32
	fmt.Printf("%v, %T\n", r, r)

}
