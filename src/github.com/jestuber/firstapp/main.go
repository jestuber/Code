package main

import (
	"fmt"
)

func main() {
	i := 10
	switch {
	case i <= 10:
		fmt.Println("less than or equal to ten")
		fallthrough // fallthrough is not implicit, breaks are implicit
	case i <= 20: // will execute next case even if it's false
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than twenty")
	}
}
