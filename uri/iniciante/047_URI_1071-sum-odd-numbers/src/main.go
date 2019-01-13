package main

import 	"bufio"
import 	"fmt"
import 	"os"
import 	"strconv"
import 	"strings"

type Input struct {
	num1 int
	num2 int
	result int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var (
		I Input
		err error
	)
	I.num1  , err = strconv.Atoi(ReadLine(reader))
	Except(err)

	I.num2  , err = strconv.Atoi(ReadLine(reader))
	Except(err)

	if (I.num1 > I.num2){
		I.num1, I.num2 = I.num2, I.num1
	}

	// fmt.Printf("de %d ate: %d ==> %d\n", I.num1, I.num2, I.num2 - I.num1)

	for index := I.num1 + 1 ; index < I.num2 ; index++ {
		if index % 2 != 0 {
			// fmt.Printf("%d + %d = %d\n", I.result, index, I.result + index)
			I.result += index
		}
	}

	fmt.Println(I.result)
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
