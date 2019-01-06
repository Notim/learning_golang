package main

import "fmt"

var distancia int
var gasto float64

func main() {
	fmt.Scanf("%d\n", &distancia)
	fmt.Scanf("%f", &gasto)
	media := float64(distancia) / gasto

	fmt.Printf("%.3f km/l\n", media)
}
