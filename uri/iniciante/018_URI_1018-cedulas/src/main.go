package main

import "fmt"

func main() {
	var valor, cont, resto int
	moedas := []int{100, 50, 20, 10, 5, 2, 1}

	fmt.Scanf("%d\n", &valor)
	fmt.Print(valor, "\n")

	for i, valorFinal := 0, valor; i < len(moedas); i++ {
		cont, resto = conversorMoeda(valorFinal, moedas[i])
		fmt.Printf("%d nota(s) de R$ %d,00\n", cont, moedas[i])
		valorFinal = resto
	}
}

func conversorMoeda(valorEntrada, moderador int) (contador, resto int) {
	resto = (valorEntrada % moderador)
	contador = (valorEntrada - resto) / moderador

	return
}
