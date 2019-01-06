/*
	para o triangulo existir precisa de que qualquer um dos lados seja menor que a soma dos outros
	e menor que a subtracao dos outros
*/

package main

import "fmt"

func main() {
	var A, B, C, rec float64
	var word string

	fmt.Scanf("%f %f %f", &A, &B, &C)

	if (A > (B-C)*1) && (A < B+C) && (B > (A-C)*1) && (B < A+C) && (C > (A-B)*1) && (C < A+B) {
		word = "Perimetro = "
		rec = A + B + C //perimetro

	} else { //senao escreva area do trapezio
		word = "Area = "
		rec = ((A + B) * C) / 2 //area
	}
	fmt.Printf("%s%.1f\n", word, rec)

}
