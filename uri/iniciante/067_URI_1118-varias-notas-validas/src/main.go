package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	reader *bufio.Reader
	err error
	notas []float64
)

func init() {
	reader = bufio.NewReader(os.Stdin)
}

func main() {
	GetNotes()
	ShowNotes()
	Ask()
}

func GetNotes()  {
	for ; len(notas) !=2 ; {
		in := ReadLine()

		num , err := strconv.ParseFloat(in, 64)
		Except(err)

		if num <= 10.0 && num >= 0.0 {
			notas = append(notas, num)
		} else {
			fmt.Println("nota invalida")
		}
	}
}

func ShowNotes()  {
	va := 0.0
	for _, el := range(notas){
		va += el
	}
	fmt.Printf("media = %.2f\n", va/2)
	notas = make([]float64, 0)
}

func Ask()  {
	line, err := strconv.Atoi(ReadLine("novo calculo (1-sim 2-nao)\n"))
	Except(err)

	switch line {
		case 1 : {
			main()
		}
		case 2: {
			os.Exit(0x00)
		}
		default: {
			Ask()
		}
	}
}

func ReadLine(msg ...string) string {
	if len(msg) > 0x00 {
		for _, e := range(msg){
			fmt.Print(e)
		}
	}

	input, err := reader.ReadString('\n')
	Except(err)

	return ClearStr(input)
}
func ClearStr(str string) string {
	str = strings.TrimSuffix(str, "\n")
	str = strings.TrimSuffix(str, "\r")
	str = strings.TrimSuffix(str, "\r\n")

	return str
}
func Except(err error){
	if err != nil {
		fmt.Println(err)
		os.Exit(0x01)
	}
}
/*
-3.5
3.5
11.0
10.0
4
1
8.0
9.0
2*/