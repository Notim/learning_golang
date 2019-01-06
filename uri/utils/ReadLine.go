package utils

import (
	"bufio"
	"os"
	"strings"
)

func ReadLine() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	return strings.Replace(input, "\n", "", -1)
}
