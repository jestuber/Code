package main

import (
	"fmt"
)

const (
	x = iota
	y // compiler infers that we want these to be iota
	z
)

const (
	x2 = iota //iota is scoped to a const block
)

func main() {
	const myConst int = 42
	fmt.Printf("%v, %T\n", myConst, myConst)

	var b int = 27
	fmt.Printf("%v, %T\n", myConst+b, myConst+b)

	const a = 42 // untyped constant
	var c int16 = 27
	fmt.Printf("%v, %T\n", a+c, a+c)
	var d int32 = 27 // this works just fine
	fmt.Printf("%v, %T\n", a+d, a+d)

	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%v\n", z)
	fmt.Printf("%v\n", x2)

}
