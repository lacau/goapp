package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n bool = true
	fmt.Printf("value: %v, type: %T\n", n, n)

	var bo bool
	fmt.Println("Variables always init with 0 value when not assigned")
	fmt.Printf("value: %v, type: %T\n", bo, bo)

	var i int = 10
	fmt.Println("int has unspecified size, every platform can be different")
	fmt.Printf("value: %v, type: %T\n", i, i)

	var i2 uint = 42
	var i3 uint8 = 42
	var i4 uint16 = 42
	var i5 uint32 = 42
	fmt.Printf("value: %v, type: %T\n", i2, i2)
	fmt.Printf("value: %v, type: %T\n", i3, i3)
	fmt.Printf("value: %v, type: %T\n", i4, i4)
	fmt.Printf("value: %v, type: %T\n", i5, i5)

	var s1 = 10
	var s2 = 3
	var r = 10 / 3
	fmt.Println("Types can not change on arithmetic operations ")
	fmt.Printf("value: %v, type: %T\n", s1, s1)
	fmt.Printf("value: %v, type: %T\n", s2, s2)
	fmt.Printf("10 / 3 = %v\n", r)

	fmt.Println("Logic binary operators")
	var a int64 = 10
	var b int64 = 3
	br1 := a & b
	br2 := a | b
	br3 := a ^ b
	br4 := a &^ b
	fmt.Printf("a = %v // %v\n", a, strconv.FormatInt(a, 2))
	fmt.Printf("b = %v // %v\n", b, strconv.FormatInt(b, 2))
	fmt.Printf("a & b = %v // %v\n", br1, strconv.FormatInt(br1, 2))
	fmt.Printf("a | b = %v // %v\n", br2, strconv.FormatInt(br2, 2))
	fmt.Printf("a ^ b = %v // %v\n", br3, strconv.FormatInt(br3, 2))
	fmt.Printf("a &^ b = %v // %v\n", br4, strconv.FormatInt(br4, 2))

	fmt.Println("Bit shit operations")
	bs1 := 8
	fmt.Println("8 // 2^3")
	fmt.Printf("%v << 3 = %v // 2^3 * 2^3 = 2^6\n", bs1, bs1<<3)
	fmt.Printf("%v >> 3 = %v // 2^3 / 2^3 = 2^0\n", bs1, bs1>>3)
}
