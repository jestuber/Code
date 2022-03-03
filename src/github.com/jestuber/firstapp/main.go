package main

import (
	"fmt"
)

func main() {
	var ms *myStruct
	ms = new(myStruct)
	(*ms).foo = 42 // * has lower precedence than . operator
	fmt.Println((*ms).foo)
	ms.foo = 27
	fmt.Println(ms.foo) // syntactic sugar to scare the c++ devs

}

type myStruct struct {
	foo int
}
