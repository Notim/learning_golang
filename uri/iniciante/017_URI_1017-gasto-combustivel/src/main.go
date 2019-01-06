package main

import "fmt"

const gastoKml = 12

var DS, DT, VM, gastoTotal float64

func main() {
	//VM = DS / DT
	fmt.Scanf("%f\n %f", &DT, &VM)

	DS = VM * DT
	gastoTotal = DS / gastoKml

	fmt.Printf("%.3f\n", gastoTotal)

}
