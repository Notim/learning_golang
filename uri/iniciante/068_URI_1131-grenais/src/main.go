package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Grenal struct {
	interWon  int
	gremioWon int
	draw  	  int
	games [][]int
}
var (
	reader *bufio.Reader
	err error
	gols []int
)

func init() {
	reader = bufio.NewReader(os.Stdin)
}
func main() {
	var Grenal Grenal

	for {
		gols = make([]int, 0)

		line := ReadLine()
		golStr := strings.Split(line, " ")

		n1, _ := strconv.Atoi(golStr[0])
		gols   = append(gols, n1)
		n2, _ := strconv.Atoi(golStr[1])
		gols   = append(gols, n2)

		Grenal.games = append(Grenal.games, gols)

		// fmt.Println(Grenal)

		in := ReadLine("Novo grenal (1-sim 2-nao)\n")
		if in != "1" {
			break
		}
	}
	fmt.Printf("%d grenais\n", len(Grenal.games))

	for _, el := range(Grenal.games) {
		if el[0] > el[1]{
			Grenal.interWon ++
		} else if el[0] < el[1]{
			Grenal.gremioWon++
		}else if el[0] == el[1]{
			Grenal.draw++
		}
	}
	fmt.Printf("Inter:%d\n", Grenal.interWon)
	fmt.Printf("Gremio:%d\n", Grenal.gremioWon)
	fmt.Printf("Empates:%d\n", Grenal.draw)

	if  Grenal.interWon >  Grenal.gremioWon {
		fmt.Print("Inter venceu mais\n")
	} else if Grenal.interWon < Grenal.gremioWon{
		fmt.Print("Gremio venceu mais\n")
	} else {
		fmt.Print("Nao houve vencedor\n")
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