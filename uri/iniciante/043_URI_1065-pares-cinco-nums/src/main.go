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
		valores [5]int
		pares []int
	)

	for index := 0 ; index < 5; index++ {
		valores [index] , _ = strconv.Atoi(ReadLine(reader))

		if ((valores[index] % 2) == 0) {
			pares = append(pares, valores[index])
		}
	}
	fmt.Printf("%d valores pares\n", len(pares))
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