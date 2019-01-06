package main

import "fmt"

func main() {
	var vector [5]float64
	var position, qtd int

	vector[0] = 4.0
	vector[1] = 4.5
	vector[2] = 5.0
	vector[3] = 2.0
	vector[4] = 1.5

	fmt.Scanf("%d", &position)
	fmt.Scanf("%d", &qtd)

	total := float64(qtd) * vector[position-1]
	fmt.Printf("Total: R$ %.2f\n", total)
}
