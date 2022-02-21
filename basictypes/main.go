package main

import (
	"fmt"
)

// Declaration's are 4 types var, const, type and func
const boilingF = 212.0 //is a package level declaration as is main

//type declaration
type Celcius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celcius = -273.15
	FreezingC     Celcius = 0
	BoilingC      Celcius = 100
)

func main() {
	//short variable declarations used for intializing local variable
	v := 0.0
	fmt.Printf("Type of v %T", v)

	variableDeclaration()
	fmt.Printf("\n CToF 30 =%v", CToF(30))
	fmt.Printf("\n FToC 32 =%v", FToC(32))
}

//type declaration demo
func CToF(c Celcius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celcius {
	return Celcius((f - 32) * 5 / 9)
}

func variableDeclaration() {
	var f = boilingF //local variable scopes within main
	var c = (f - 32) / 5 * 9
	fmt.Printf("Boiling point =%fF or %gC\n", f, c)

	//varible declaration
	var d, e, g = 1, 2.0, "Hello"
	fmt.Printf("d= %v, e= %v, f= %v", d, e, g)
}

func pointerVars() {
	//Pointers
	fmt.Println("Pointer in Go!!")
	//var p1 = gi()
	fmt.Println(gi() == gi())

}
func gi() *int {
	z := 1
	return &z
}
