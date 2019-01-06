package main

import (
	"fmt"
)

func main() {
	//area = π . raio2
	const _PI = 3.14159
	var area, raio float64
	fmt.Scanf("%f\n", &raio)
	area = ((raio * raio) * _PI)
	fmt.Printf("A=%.4f\n", area)
}
