package main

import (
	"fmt"
	"strconv"
)

var i int = 27

func main() {
	var i int = 42
	fmt.Printf("%v, %T\n", i, i)

	var j string
	j = strconv.Itoa(i) // Integer to ASCII
	fmt.Printf("%v, %T\n", j, j)
}
