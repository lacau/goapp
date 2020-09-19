package main

import "fmt"

func main() {
	statePopulation := make(map[string]int)
	statePopulation = map[string]int{
		"California": 39250017,
		"Texas":      27862596,
		"Florida":    20612439,
		"New York":   19745289,
	}

	fmt.Println(statePopulation)
	fmt.Println("PS: arrays can be map's key, but slices can't!")
	fmt.Println("WARNING: maps are initialized as references")
	fmt.Printf("key: Texas, value: %v\n", statePopulation["Texas"])

	statePopulation["Test"] = 55000
	fmt.Println(statePopulation)
	delete(statePopulation, "Test")
	fmt.Println(statePopulation)

	pop, ok := statePopulation["Test"]
	fmt.Println(pop, ok)
	pop, ok = statePopulation["Florida"]
	fmt.Println(pop, ok)
}
