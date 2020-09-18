package main

import "fmt"

func main() {
	grades := []int{97, 85, 93}
	fmt.Printf("Grades: %v\n", grades)

	var students [3]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Leandro"
	students[1] = "Alinne"
	students[2] = "Sophia"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Number of students: %v\n", len(students))

	fmt.Println("WARING: When assign an array to another variable and it was declared with [...], it is a copy!")

	a := [...]int{1, 2, 3}
	b := a
	a[2] = 5
	fmt.Println(a)
	fmt.Println(b)
}
