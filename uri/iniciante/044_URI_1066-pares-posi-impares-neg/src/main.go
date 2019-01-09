package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// @struct with counters and Values
type Numbers struct {
	oddCount int 	// impar
	evenCount int 	// par
	positives int
	negatives int

	values []int
}

// @method of NumClassification like a static method

func (numInfo *Numbers) PutValue(value int) {
	numInfo.values = append(numInfo.values, value)
}

func (numInfo *Numbers) BuildStats() {
	for _, element := range(numInfo.values) {
		if element > 0 {
			numInfo.positives ++
		}
		if element < 0 {
			numInfo.negatives ++
		}
		if element % 2 == 0 {
			numInfo.evenCount ++
		}
		if element % 2 != 0 {
			numInfo.oddCount ++
		}
	}
}

// @method of NumClassification
func (numInfo Numbers) ShowClassification(lang string) Numbers{
	if IsEmpty(lang) { lang = "en-us" }

	switch  {
		case lang == "pt-br" :
			fmt.Printf("%d valor(es) par(es)\n"		, numInfo.evenCount)
			fmt.Printf("%d valor(es) impar(es)\n"	, numInfo.oddCount)
			fmt.Printf("%d valor(es) positivo(s)\n"	, numInfo.positives)
			fmt.Printf("%d valor(es) negativo(s)\n"	, numInfo.negatives)

		case lang == "en-us" :
			fmt.Printf("%d even Values\n"		, numInfo.evenCount)
			fmt.Printf("%d odd Values\n"		, numInfo.oddCount)
			fmt.Printf("%d positive Values\n"	, numInfo.positives)
			fmt.Printf("%d Negative Values\n"	, numInfo.negatives)
		default:
	}

	return numInfo
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var (
		info Numbers
		input int
	)

	for index := 0 ; index < 5; index++ {
		input, _ = strconv.Atoi(ReadLine(reader))

		info.PutValue(input)
	}

	info.BuildStats()
	info.ShowClassification("pt-br") // like a C# extension method

	// info = NumClassification.BuildNumClassification(info) // like a C# static method
	// info.ShowClassification("pt-br") // single call extension mode
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