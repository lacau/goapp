package main

import "fmt"

func main() {
	n := 3.14
	n = 2.1e2
	fmt.Printf("%v, %T\n", n, n)

	fmt.Println("Complex numbers")
	fmt.Println("complex64 = float32 + float32")
	fmt.Println("complex128 = float64 + float64")
	var c complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("real = %v, %T\n", real(c), real(c))
	fmt.Printf("imag = %v, %T\n", imag(c), imag(c))
}
