package main

import 	"bufio"
import 	"fmt"
import 	"os"
import 	"strconv"
import 	"strings"

type ListInt []int

type numbers struct {
	inCount int
	outCount int
	values ListInt
}

func (T *numbers) Add(value int){
	T.values = append(T.values, value)
}

func (T ListInt) Lenght() int{
	return len(T)
}

func (T *ListInt) Add(value int) {
	*T = append(*T, value)
}

// predicate func(TModel *numbers) bool
func (T numbers) Filter(predicate func(current int) bool) (tempList ListInt){
	for index := 0; index < len(T.values); index++ {
		if predicate(T.values[index]) {
			tempList = append(tempList, T.values[index])
		}
	}

	return
}
func Find(current int) (boo bool) {
	return current > 0
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

	// Filter(i => i > 0 ) ==> lambda epr
	num.inCount  = num.Filter(func(i int) bool {
		return i >= 10 && i <= 20
	}).Lenght()

	num.outCount = num.Filter(func(i int) bool {
		return i < 10 || i > 20
	}).Lenght()

	fmt.Printf("%v in\n", num.inCount)
	fmt.Printf("%v out\n", num.outCount)
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
