package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := a
	b[2] = 5
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
	fmt.Println(b)
	fmt.Printf("Length: %v\n", len(b))
	fmt.Printf("Capacity: %v\n", cap(b))

	fmt.Println("Slice syntax x[y:z] - y is inclusive and z is exclusive")
	fmt.Println("both params are optional so, x[:] is a valid syntax")
	fmt.Println("WARNING: Slices initialize as a pointer by default")

	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s2 := s1[:]
	s3 := s1[3:]
	s4 := s1[:6]
	s5 := s1[3:6]
	fmt.Printf("%v\n", s1)
	fmt.Printf("%v\n", s2)
	fmt.Printf("%v\n", s3)
	fmt.Printf("%v\n", s4)
	fmt.Printf("%v\n", s5)

	fmt.Println("Using make to create slices")
	m1 := make([]int, 3, 100)
	fmt.Println(m1)
	fmt.Printf("Length: %v\n", len(m1))
	fmt.Printf("Capacity: %v\n", cap(m1))

	si := []int{}
	fmt.Println(si)
	fmt.Printf("Length: %v\n", len(si))
	fmt.Printf("Capacity: %v\n", cap(si))

	si = append(si, 1)
	fmt.Println(si)
	fmt.Printf("Length: %v\n", len(si))
	fmt.Printf("Capacity: %v\n", cap(si))

	si = append(a, []int{1, 2, 3, 4, 5}...)
	fmt.Println(si)
	fmt.Printf("Length: %v\n", len(si))
	fmt.Printf("Capacity: %v\n", cap(si))
}
