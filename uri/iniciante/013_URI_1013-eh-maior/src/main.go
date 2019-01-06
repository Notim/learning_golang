package main

import "fmt"

var vector [3]int
var maior, index int

func main() {
	for index = 0; index < 3; index++ {
		fmt.Scanf("%d", &vector[index])
		if vector[index] > maior {
			maior = vector[index]
		}
	}
	fmt.Printf("%v eh o maior\n", maior)
}
