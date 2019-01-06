package main

import "fmt"
import "strings"
import "strconv"

func main() {
	var cont, resto int

	var valorS string

	cedulas := []int{100, 50, 20, 10, 5, 2, 1}
	moedas := []int{50, 25, 10, 5, 1}

	fmt.Scanf("%s\n", &valorS)
	var rec []string = strings.Split(valorS, ".")

	valorAntesComma, err := strconv.Atoi(rec[0])
	valorDecimal, err := strconv.Atoi(rec[1])
	if !(err == nil) {
		fmt.Print("Erro")
	}
	fmt.Print("NOTAS:\n")
	for i, valorFinal := 0, valorAntesComma; i < len(cedulas); i++ {
		cont, resto = conversorMoeda(valorFinal, cedulas[i])
		if i == (len(cedulas) - 1) {
			fmt.Print("MOEDAS:\n")
			fmt.Printf("%d moeda(s) de R$ %d.00\n", cont, cedulas[i])

		} else {
			fmt.Printf("%d nota(s) de R$ %d.00\n", cont, cedulas[i])
		}
		valorFinal = resto
	}
	for i, valorFinal := 0, valorDecimal; i < len(moedas); i++ {
		cont, resto = conversorMoeda(valorFinal, moedas[i])
		fmt.Printf("%d moeda(s) de R$ 0.%02d\n", cont, moedas[i])
		valorFinal = resto
	}
}
func conversorMoeda(valorEntrada, moderador int) (contador, resto int) {
	resto = (valorEntrada % moderador)
	contador = (valorEntrada - resto) / moderador

	return
}
