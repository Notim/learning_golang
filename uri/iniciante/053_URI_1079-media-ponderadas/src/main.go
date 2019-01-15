package main

import (
	"bufio"
)
import 	"fmt"
import 	"os"
import 	"strconv"
import 	"strings"

type Notas struct {
	nota1 float64
	nota2 float64
	nota3 float64
}
type ListaNotas []Notas

func (T *ListaNotas) Add(value Notas) {
	*T = append(*T, value)
}

func (T *ListaNotas) Foreach(action func(element Notas)) {
	for _, element := range(*T) {
		action(element)
	}
}

func (T *Notas) GetMedia(p1, p2, p3 int) (result float64) {
	result = ((T.nota1 * float64(p1)) + (T.nota2 * float64(p2)) + (T.nota3 * float64(p3))) / float64(p1 + p2 + p3)

	return
}

func main() {
	var err error
	var notas ListaNotas

	reader := bufio.NewReader(os.Stdin)

	loop, err := strconv.Atoi(ReadLine(reader))
	Except(err)

	for index := 1; index <= loop; index++ {
		sentence := ReadLine(reader)
		val := strings.Split(sentence, " ")

		no1, err := strconv.ParseFloat(val[0], 64)
		Except(err)

		no2, err := strconv.ParseFloat(val[1], 64)
		Except(err)

		no3, err := strconv.ParseFloat(val[2], 64)
		Except(err)

		notas.Add(Notas{
			nota1 : no1,
			nota2 : no2,
			nota3 : no3,
		})
	}
	notas.Foreach(func(x Notas) {
		fmt.Printf("%.1f\n", x.GetMedia(2, 3, 5))
	})
}

func ReadLine(reader *bufio.Reader) string {
	input, err := reader.ReadString('\n')
	Except(err)

	return ClearStr(input)
}

func ClearStr(str string) string {
	str = strings.TrimSuffix(str, "\n")
	str = strings.TrimSuffix(str, "\r")
	str = strings.TrimSuffix(str, "\t")
	str = strings.TrimSuffix(str, "\r\n")

	return str
}

func Except(err error) {
	if (err != nil) {
		fmt.Println(err)
		os.Exit(1)
	}
}