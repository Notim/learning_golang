package main

import "bufio"
import 	"fmt"
import 	"os"
import 	"strconv"
import 	"strings"
import . "utils/Enumerables/List"

type structExp struct {
	rabbit int
	rat int
	frog int
	total int
	percentage map[string]float64
	listString ListString
}

func (T *structExp) GetTotal() {
	T.total = T.rat + T.frog + T.rabbit
}

func (T *structExp) Percentage() {
	T.GetTotal()

	T.percentage = make(map[string]float64) // allocate memory to the map (instance)

	T.percentage["rabbit"] = (float64(T.rabbit) / float64(T.total) * 100)
	T.percentage["frog"]   = (float64(T.frog) 	/ float64(T.total) * 100)
	T.percentage["rat"]	   = (float64(T.rat) 	/ float64(T.total) * 100)
}

func (T *structExp) Present() {
	fmt.Printf("Total: %d cobaias\n", T.total)
	fmt.Printf("Total de coelhos: %d\n", T.rabbit)
	fmt.Printf("Total de ratos: %d\n", T.rat)
	fmt.Printf("Total de sapos: %d\n", T.frog)
	fmt.Printf("Percentual de coelhos: %.2f %%\n", T.percentage["rabbit"])
	fmt.Printf("Percentual de ratos: %.2f %%\n", T.percentage["rat"])
	fmt.Printf("Percentual de sapos: %.2f %%\n", T.percentage["frog"])
	T.listString.Add("Hello Muchacho !!!")
	fmt.Println(T)
}

func main() {
	var loop int
	var Exp structExp

	reader := bufio.NewReader(os.Stdin)

	loop , err := strconv.Atoi(ReadLine(reader))
	Except(err)

	for index := 0; index != loop; index++ {
		sentence := ReadLine(reader)
		Except(err)

		value := strings.Split(sentence, " ")

		val ,  err := strconv.Atoi(value[0])
		Except(err)

		switch {
			case value[1] == "C":
				Exp.rabbit += val
				break

			case value[1] == "S":
				Exp.frog += val
				break

			case value[1] == "R":
				Exp.rat += val
				break
			default:
				fmt.Println("out", value[1])
		}
	}

	Exp.GetTotal()
	Exp.Percentage()
	Exp.Present()
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