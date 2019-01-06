package main

import "fmt"

var nome string
var salarioFixo, vendas, salarioTotal float64

func main() {
	fmt.Scanf("%s\n %f\n %f\n", &nome, &salarioFixo, &vendas)
	salarioTotal = float64(salarioFixo + (vendas * 0.15))
	fmt.Printf("TOTAL = R$ %.2f\n", salarioTotal)
}
