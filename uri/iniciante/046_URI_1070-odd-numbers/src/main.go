package main

import 	"bufio"
import 	"fmt"
import 	"os"
import 	"strconv"
import 	"strings"

func main() {
	reader := bufio.NewReader(os.Stdin)
	var (
		input int
	)

	input, _ = strconv.Atoi(ReadLine(reader))

	for index := input; index != (input + 12); index++ {
		if index % 2 != 0 {
			fmt.Println(index)
		}
	}
}


// Util functions
func ReadLine(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')

	return ClearStr(input)
}

func ClearStr(str string) string {
	str = strings.TrimSuffix(str, "\n")
	str = strings.TrimSuffix(str, "\r")
	str = strings.TrimSuffix(str, "\t")
	str = strings.TrimSuffix(str, "\r\n")

	return str
}

func IsEmpty(exp string) bool {
	return len(exp) == 0
}