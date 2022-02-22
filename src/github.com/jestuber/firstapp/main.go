package main

import "fmt"

func main() {
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"California": 39250017,
		"Texas":      27862596,
		"Ohio":       11614373,
	}
	pop, ok := statePopulations["typo"] // use comma ok to check if value is 0 or not in map
	fmt.Println(pop, ok)
	_, ok = statePopulations["Ohio"]
	fmt.Println(ok)
}
