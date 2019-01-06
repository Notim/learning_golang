package main

import (
	"fmt"
)

func main() {
	var notaA, notaB, notaC float64

	fmt.Scanf("%f\n", &notaA)
	fmt.Scanf("%f\n", &notaB)
	fmt.Scanf("%f\n", &notaC)

	resultado := ((notaA * 2) + (notaB * 3) + (notaC * 5)) / 10

	fmt.Printf("MEDIA = %.1f\n", resultado)
}
