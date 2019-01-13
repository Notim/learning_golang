package main

import (
	"bufio"
	"math"
)
import 	"fmt"
import 	"os"
import 	"strconv"
import 	"strings"

func main() {
	reader := bufio.NewReader(os.Stdin)

	var (
		err error
	)

	loop , err := strconv.Atoi(ReadLine(reader))
	Except(err)

	for index := 1; index <= loop; index++ {
		if index % 2 == 0 {
			fmt.Println(math.Pow(float64(index), 2))
		}
	}
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