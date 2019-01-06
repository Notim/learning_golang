package main

import io "fmt"
import "os"

func main() {
	var value [3]float64
	var situacao [2]float64
	var swap, A, B, C float64

	/*
		resultados para situacao situacao
		situacao[0] = 0 => "NAO FORMA TRIANGULO"
		situacao[0] = 1 => "TRIANGULO RETANGULO"
		situacao[0] = 2 => "TRIANGULO OBTUSANGULO"
		situacao[0] = 3 => "TRIANGULO ACUTANGULO"

		situacao[1]= 0 => null
		situacao[1]= 1 => "TRIANGULO EQUILATERO"
		situacao[1]= 2 => "TRIANGULO ISOSCELES"
	*/

	io.Scanf("%f", &value[0]) //6 36
	io.Scanf("%f", &value[1]) //8 64
	io.Scanf("%f", &value[2]) //10 100

	//bubble sort
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			if value[j] > value[j+1] {
				swap = value[j]
				value[j] = value[j+1]
				value[j+1] = swap
			}
		}
	}

	A = value[2]
	B = value[1]
	C = value[0]

	situacao[1] = 0
	if A >= B+C {
		situacao[0] = 0
	} else if (A * A) == (B*B + C*C) {
		situacao[0] = 1
	} else if (A * A) > (B*B + C*C) {
		situacao[0] = 2
	} else if (A * A) < (B*B + C*C) {
		situacao[0] = 3
	}

	if (A == B) || (A == C) || (B == C) {
		situacao[1] = situacao[1] + 1
	}
	if (A == C) && (B == C) {
		situacao[1] = situacao[1] + 1
	}

	if situacao[0] == 0 {
		io.Printf("NAO FORMA TRIANGULO\n")
	} else if situacao[0] == 1 {
		io.Printf("TRIANGULO RETANGULO\n")
	} else if situacao[0] == 2 {
		io.Printf("TRIANGULO OBTUSANGULO\n")
	} else if situacao[0] == 3 {
		io.Printf("TRIANGULO ACUTANGULO\n")
	}

	if situacao[1] == 2 {
		io.Print("TRIANGULO EQUILATERO\n")
	} else if situacao[1] == 1 {
		io.Printf("TRIANGULO ISOSCELES\n")
	}

	os.Exit(0)
}
