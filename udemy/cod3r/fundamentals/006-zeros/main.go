package main

import "fmt"

func main() {
	var (
		zeroInt    int
		zeroFloat  float64
		zeroString string
		zeroBool   bool
	)
	fmt.Println()
	fmt.Printf("zero value of int64 is: %d\n", zeroInt)
	fmt.Printf("zero value of float64 is: %f\n", zeroFloat)
	fmt.Printf("zero value of string is: %s\n", zeroString)
	fmt.Printf("zero value of bool is: %t\n", zeroBool)

}
