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

	for  {
		loop := ReadLine()

		nums = strings.Split(loop, " ")

		n1, err := strconv.Atoi(nums[0x00])
		Except(err)

		n2 , err := strconv.Atoi(nums[0x01])
		Except(err)

		if (n1 == 0x00 || n2 == 0x00){
			break
		}

		pair = append(pair, []int{n1, n2})
	}
	/*  x  y ==> primeiro
	   -x  y ==> segundo
	   -x -y ==> terceiro
	    x -y ==> quarto
 	    0  2 ==> end
	*/
	for _, e := range (pair) {
		switch {
		case e[0] > 0 : {
			if e[1] > 0 {
				fmt.Println("primeiro")
			} else {
				fmt.Println("quarto")
			}
		}
		case e[0] < 0 : {
			if e[1] > 0 {
				fmt.Println("segundo")
			} else {
				fmt.Println("terceiro")
			}
		}
		default:{
			fmt.Println("nem era pra entrar aqui")
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