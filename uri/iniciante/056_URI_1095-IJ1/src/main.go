package main

import "fmt"

func main() {
	for i, j := 1, 60; j >= 0 ; i = i + 3 {
		fmt.Printf("I=%d J=%d\n",i, j)
		j = j - 5
	}
}
