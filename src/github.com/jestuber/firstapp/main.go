package main

import (
	"fmt"
)

func main() {
	var a int = 42
	var b *int = &a // pointer
	fmt.Println(a, b)
	fmt.Println(&a, b)
	fmt.Println(a, *b) // dereference
	a = 27
	fmt.Println(a, *b)
	*b = 14
	fmt.Println(a, *b)

}
