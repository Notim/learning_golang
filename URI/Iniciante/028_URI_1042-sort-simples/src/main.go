package main

import "fmt"

func main() {
	var slice []float64
	var in float64

	for i := 0; i < 3; i++ {
		fmt.Scanf("%f", &in)
		slice = append(slice, in)
	}

	sorted := ReturnSorted(slice)

	PrintSlice(sorted)
	fmt.Print("\n")
	PrintSlice(slice)
}

func ReturnSorted(slice []float64) []float64 {
	var swap, valor float64
	var internSlice []float64
	var slicepointer *float64

	for i := 0; i < len(slice); i++ {
		slicepointer = &slice[i]
		valor = *slicepointer
		internSlice = append(internSlice, valor)
	}

	for i := 0; i < len(internSlice); i++ {
		for j := 0; j < (len(internSlice) - 1); j++ {
			if internSlice[j] > internSlice[j+1] {
				swap = internSlice[j]
				internSlice[j] = internSlice[j+1]
				internSlice[j+1] = swap
			}
		}
	}
	return internSlice
}

func PrintSlice(slice []float64) {
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
}
