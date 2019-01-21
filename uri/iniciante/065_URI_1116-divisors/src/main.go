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
	pair [][]float64
	nums []string
)

func init() {
	reader = bufio.NewReader(os.Stdin)
	nums = []string{"0", "0"}
}

func main() {
	loop, er := strconv.Atoi(ReadLine())
	Except(er)

	for index := 0x00; index < loop; index++ {
		line := ReadLine()

		nums = strings.Split(line, " ")

		num1, err := strconv.ParseFloat(nums[0x00], 64)
		Except(err)

		num2, err := strconv.ParseFloat(nums[0x01], 64)
		Except(err)

		pair = append(pair, []float64{num1, num2})
	}

	for _, e := range (pair) {
		dix := e[0] / e[1]

		if (e[1] == 0){
			fmt.Printf("divisao impossivel\n")
		}else {
			fmt.Printf("%.1f\n", dix)
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