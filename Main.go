package main

import (
	"fmt"
)

const (
	a = iota
	b
	c
)

const (
	a2 = iota
)

const (
	_ = iota
	firstValue
)

func main() {
	fmt.Println("Working with constants")

	const myConst1 int = 42
	fmt.Printf("%v, %T\n", myConst1, myConst1)

	const myConst2 float64 = 42.42
	fmt.Printf("%v, %T\n", myConst2, myConst2)

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", a2)
	fmt.Printf("%v\n", firstValue)

	fmt.Println("primitive types can be constant, except arrays")
	fmt.Println("arithmetic operations can be perform between constants and variables")
	fmt.Println("WARNING: package and global scope constants be shadowed")
	fmt.Println("use _ on enumerated constants to drop allocated value")
}
