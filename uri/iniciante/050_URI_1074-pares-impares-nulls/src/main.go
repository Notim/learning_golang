package main

import 	"bufio"
import 	"fmt"
import 	"os"
import 	"strconv"
import 	"strings"

type ListInt []int

type numbers struct {
	values ListInt
}
func (T ListInt) Lenght() int{
	return len(T)
}
func (T *ListInt) Add(value int) {
	*T = append(*T, value)
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

		num.values.Add(in)
	}

	for _ , value := range(num.values) {
		switch {
			case value > 0 && value % 2 == 0:
				fmt.Printf("EVEN POSITIVE\n")
				break

			case value < 0 && value % 2 == 0:
				fmt.Printf("EVEN NEGATIVE\n")
				break

			case value > 0 && value % 2 != 0:
				fmt.Printf("ODD POSITIVE\n")
				break

			case value < 0 && value % 2 != 0:
				fmt.Printf("ODD NEGATIVE\n")
				break
			default:
				fmt.Printf("NULL\n")
				break
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