package main

import (
	"fmt"
)

func main() {
	a := "start"
	defer fmt.Println(a) // defer takes arguments at the time defer is called
	a = "end"
}
