package main

import "fmt"

func main() {
	var valorEntrada, hora, minuto, segundo, resto int

	fmt.Scanf("%d\n", &valorEntrada)

	resto = valorEntrada % 3600
	hora = (valorEntrada - resto) / 3600
	valorEntrada = resto

	resto = valorEntrada % 60
	minuto = (valorEntrada - resto) / 60
	segundo = resto

	fmt.Print(hora, ":", minuto, ":", segundo, "\n")
}
