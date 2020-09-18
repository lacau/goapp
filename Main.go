package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 30.5
	fmt.Printf("value: %v, type: %T\n", a, a)
	fmt.Printf("value: %v, type: %T\n", b, b)

	a = int(b)
	fmt.Printf("value: %v, type: %T\n", a, a)
}
