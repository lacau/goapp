package main

import "fmt"

func main() {
	s := "This is a string"
	fmt.Printf("%v, %T\n", s, s)
	fmt.Printf("%v, %T\n", s[2], s[2])
	fmt.Println("string are aliases for bytes")
	fmt.Printf("%v, %T\n", string(s[2]), s[2])
	fmt.Println("strings are generally immutable")

	fmt.Printf("s = %v\n", s)
	fmt.Printf("s + s = %v\n", s+s)

	fmt.Println("string to byte slices")
	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b)

	fmt.Println("runes 'x', runes are a type alias for int32")
	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r)
}
