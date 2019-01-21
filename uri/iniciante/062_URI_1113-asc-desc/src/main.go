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
	pair [][]int
	nums []string
)

func init() {
	reader = bufio.NewReader(os.Stdin)
	nums = []string{"0", "0"}
}

func main() {
	goto loop

	loop :{
		loop := ReadLine()

		nums = strings.Split(loop, " ")

		n1, err := strconv.Atoi(nums[0x00])
		Except(err)

		n2 , err := strconv.Atoi(nums[0x01])
		Except(err)

		if (n1 == n2){
			goto ordering
		}

		pair = append(pair, []int{n1, n2})

		goto loop
	}

	ordering: {
		for _, e := range (pair) {
			if e[0] > e[1] {
				fmt.Printf("Decrescente\n")
			} else {
				fmt.Printf("Crescente\n")
			}
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