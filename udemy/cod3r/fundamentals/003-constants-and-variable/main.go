package main

import (
	io "fmt"
	"math"
)

func main() {
	const PI float64 = 3.14159265359 // this is a constant
	var radius float64 = 10          // and this is a local variable explicit form

	area := math.Pow(radius, 2) * PI // this is a local variable implicit form

	var a, b, g = true, false, true // in go is possible put values in various variables in same line

	/*
	   in other languages when we need change values with a neutral variable we use a swap variable

	   like this:

	   swap = a
	   a = b
	   b = swap

	   in golang is most easy to do this:
	*/
	a, b = b, a // look this shit huhueheu

	io.Println(a, b, g)                     // Golang doesnt permit unused variables
	io.Printf("circle Area is: %f\n", area) // Golang doesnt permit unused variables
}
