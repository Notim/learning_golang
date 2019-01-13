package main

import 	"bufio"
import 	"fmt"
import 	"os"
import 	"strconv"
import 	"strings"

type numbers struct {
	inCount int
	outCount int
	values []int
}

func (T *numbers) Add(value int){
	T.values = append(T.values, value)
}


func main() {
	reader := bufio.NewReader(os.Stdin)

	var (
		num numbers
		err error
	)

	loop , err := strconv.Atoi(ReadLine(reader))
	(func(err error) {
		if err != nil{
			fmt.Println(err)
			os.Exit(1)
		}
	})(err)

	for index := 0; index < loop; index++ {
		in , err := strconv.Atoi(ReadLine(reader))
		Except(err)

		num.Add(in)
	}

	// fmt.Printf("%v", num)
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
