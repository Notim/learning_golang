package main

import "fmt"
import "math"
import "os"

func main() {
	var vector [3]float64
	var result [2]float64

	for i := 0; i < 3; i++ {
		fmt.Scan(&vector[i])
	}

	result = calculo(vector[:])
	if math.IsNaN(result[0]) || math.IsNaN(result[1]) {
		fmt.Print("Impossivel calcular\n")
		os.Exit(0)
	}

	fmt.Printf("R1 = %.5f\nR2 = %.5f\n", result[0], result[1])

}
func calculo(vector []float64) (result [2]float64) {
	a, b, c := vector[0], vector[1], vector[2]
	Delta := (math.Pow(b, 2)) - 4*(a*c)

	result[0] = ((-b) + math.Sqrt(Delta)) / (2 * a)
	result[1] = ((-b) - math.Sqrt(Delta)) / (2 * a)

	return
}
