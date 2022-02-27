package main

import (
	"fmt"
)

func main() {
	statePopulations := map[string]int{
		"California": 39250017,
		"Texas":      27862596,
	}
	if pop, ok := statePopulations["Texas"]; ok {
		fmt.Println(pop)
	}
	if pop, ok := statePopulations["Florida"]; ok {
		fmt.Println(pop) // Initializer syntax
	}
}
