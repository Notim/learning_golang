package main

import "bufio"
import 	"fmt"
import 	"os"
import 	"strconv"
import 	"strings"

type structBiggest struct {
	index int
	value int
}
func main() {
	var Big structBiggest
	var values []int

	reader := bufio.NewReader(os.Stdin)

	for index := 1; index <= 100; index++ {
		value, err := strconv.Atoi(ReadLine(reader))
		Except(err)

		values = append(values, value)
	}

	for index, element := range(values) {
		if index > 0 {
			// fmt.Println(element," ==> ",values[index - 1])
			if element > Big.value {
				Big.index, Big.value = index+1, element
			}
		}else {
			Big.index, Big.value = index+1, element
		}
	}
	fmt.Printf("%d\n%d\n",Big.value, Big.index)
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