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
)

func init() {
	reader = bufio.NewReader(os.Stdin)
}

func main() {
	goto loop

	loop :{
		pass, er := strconv.Atoi(ReadLine())
		Except(er)

		if (pass != 0x7D2){
			goto incorect
		}
		goto correct
	}
	incorect :{
		fmt.Println("Senha Invalida")
		goto loop
	}
	correct: {
		fmt.Println("Acesso Permitido")
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