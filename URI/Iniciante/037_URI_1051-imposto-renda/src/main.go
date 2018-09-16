package main

import (
	io "fmt"
	"os"
)

func main() {
	os.Exit(App())
}

func Show(somatoria float64, caso int) {
	switch caso {
	case 0:
		io.Printf("R$ %.2f\n", somatoria)
		break
	case 1:
		io.Println("Isento")
		break
	}
}
func App() int {
	var salario, salarioFinal float64

	io.Scanf("%f", &salario)

	if salario < 2000.01 {
		Show(0, 1)
	} else if salario < 3000.01 {
		salario = salario - 2000
		salarioFinal = 0.08 * salario
		Show(salarioFinal, 0)

	} else if salario < 4500.01 {
		salario = salario - 3000
		salarioFinal = 80 + (0.18 * salario)
		Show(salarioFinal, 0)

	} else if salario > 4500 {
		salario = salario - 4500
		salarioFinal = 80 + 270 + (0.28 * salario)
		Show(salarioFinal, 0)
	}

	return 0
}
