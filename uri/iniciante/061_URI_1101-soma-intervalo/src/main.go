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
	for {
		loop := ReadLine()

		nums = strings.Split(loop, " ")

		n1, err := strconv.Atoi(nums[0])
		Except(err)

		n2 , err := strconv.Atoi(nums[1])
		Except(err)

		if (n1 <= 0 || n2 <= 0){
			break
		}

		pair = append(pair, []int{n1, n2})
	}

	for _, e := range(pair) {
		// fmt.Println(e)

		if e[0] < e[1] {
			e[1], e[0] = e[0], e[1]
		}
		// fmt.Println(e[0], e[1])
		sum := 0
		for i := e[1]; i <= e[0] ; i++ {
			sum += i
			fmt.Printf("%d ", i)
		}
		fmt.Printf("Sum=%d\n", sum)
	}
	// fmt.Println(pair)
}

func ReadLine(msg ...string) string {
	if len(msg) > 0{
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
		os.Exit(1)
	}
}