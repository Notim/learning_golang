package main

import "fmt"

func main() {
	var cont int
	var in int

	for i := 0; i != 6; i++ {

		fmt.Scanf("%d", &in)

		if in > 0 {
			cont++
		}
	}
	fmt.Printf("%d valores positivos\n", cont)
}
