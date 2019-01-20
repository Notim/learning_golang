package main

import "fmt"

func main() {
	var out int = 7
	var last int

	for i := 1; i <= 9 ; i = i + 2 {
		last = out
		for index := 0; index != 3; index++{
			fmt.Printf("I=%d J=%d\n", i, out)
			out--
			last++
		}
		out = last-1
	}



}
