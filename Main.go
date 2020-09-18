package main

import (
	"fmt"
)

var GlobalScopevar string = "global scope variable"
var pkgScopeVar int = 50

func main() {
	fmt.Printf("value: %v, type: %T\n", GlobalScopevar, GlobalScopevar)
	fmt.Printf("value: %v, type: %T\n", pkgScopeVar, pkgScopeVar)

	var funcScopeVar1 float32 = 30.5
	funcScopeVar2 := 30

	fmt.Printf("value: %v, type: %T\n", funcScopeVar1, funcScopeVar1)
	fmt.Printf("value: %v, type: %T\n", funcScopeVar2, funcScopeVar2)

	GlobalScopevar = "shadowing global scope variable"
	pkgScopeVar = 49

	fmt.Println("Shadow variables")
	fmt.Printf("value: %v, type: %T\n", GlobalScopevar, GlobalScopevar)
	fmt.Printf("value: %v, type: %T\n", pkgScopeVar, pkgScopeVar)
}
