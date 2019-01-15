package main

import (
	"bufio"
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

	num , err := strconv.Atoi(ReadLine(reader))
	Except(err)

	for index := 1; index <= 10; index++ {
		fmt.Printf("%d x %d = %d\n",index, num, (index * num))
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