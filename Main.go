package main

import "fmt"

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials
)

func main() {
	var roles byte = isAdmin | isHeadquarters
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is HQ? %v\n", isHeadquarters&roles == isHeadquarters)
	fmt.Printf("Can see financials? %v\n", canSeeFinancials&roles == canSeeFinancials)
}
