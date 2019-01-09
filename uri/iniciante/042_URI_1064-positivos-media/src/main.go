package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var (
		notas [6]float64
		notas4media []float64
	)

	for index := 0 ; index < 6; index++ {
		notas[index] , _ = strconv.ParseFloat(ReadLine(reader), 64)
		if (notas[index] > 0) {
			notas4media = append(notas4media, notas[index])
		}
	}

	var media float64
	for i := 0;  i < len(notas4media); i++  {
		media += notas4media[i]
	}
	media = media / float64(len(notas4media))

	fmt.Printf("%d valores positivos\n",len(notas4media))
	fmt.Printf("%.1f\n", media)
}

func ReadLine(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')

	return ClearStr(input)
}

func ClearStr(str string) string {
	str = strings.TrimSuffix(str, "\n")
	str = strings.TrimSuffix(str, "\r")
	str = strings.TrimSuffix(str, "\t")
	str = strings.TrimSuffix(str, "\r\n")

	return str
}