package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

func init() {
	reader = bufio.NewReader(os.Stdin)
}

func main() {
	// in := ReadLine("Ola Amigo ", "Como Vai")
	loop , err := strconv.Atoi(ReadLine())
	Except(err)

	for i := 0; i < loop; i++ {
		nums := strings.Split(ReadLine(), " ")

		n1, err := strconv.Atoi(nums[0])
		Except(err)

		n2, err := strconv.Atoi(nums[1])
		Except(err)

		if (n1 < n2){
			n1, n2 = n2, n1
		}

		sum := 0
		// fmt.Println(n2, n1)
		for j := n2 + 1; j < n1; j++ {
			if j % 2 != 0 {
				sum += j
			}
		}
		fmt.Println(sum)
	}
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