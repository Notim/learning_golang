package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type DateTime struct {
	day int
	hour int
	minutes int
	seconds int
	allInSec int
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(reflect.TypeOf(reader))

	var (
		FirstDay DateTime
		SecondDay DateTime
		data string
		sentence string
	)

	data = ReadLine(reader)
	sentence = ReadLine(reader)

	FirstDay = ReadDate(sentence, data)

	data = ReadLine(reader)
	sentence = ReadLine(reader)

	SecondDay = ReadDate(sentence, data)

	Differ := CompareDates(FirstDay, SecondDay)

	fmt.Printf("%d dia(s)\n", Differ.day)
	fmt.Printf("%d hora(s)\n", Differ.hour)
	fmt.Printf("%d minuto(s)\n", Differ.minutes)
	fmt.Printf("%d segundo(s)\n", Differ.seconds)

}

func ReadDate(dataSentence string , MonthDay string) DateTime {

	var(
		d DateTime
		time []string
		day []string
	)
	// var err error

	day  = strings.Split(MonthDay, " ")
	time = strings.Split(dataSentence, " : ")

	d.day , _ = strconv.Atoi(ClearStr(day[1]))

	d.hour , _ = strconv.Atoi(ClearStr(time[0]))
	d.minutes , _ = strconv.Atoi(ClearStr(time[1]))
	d.seconds , _ = strconv.Atoi(ClearStr(time[2]))
	// fmt.Println(err)

	d.allInSec = (d.hour * 3600) + (d.minutes * 60) + (d.seconds)

	return d
}

func CompareDates(d1, d2 DateTime) (ret DateTime) {
	var totalSeg  = ((d2.allInSec - d1.allInSec) + ((d2.day - d1.day) * 86400)) * 1

	Sdias 		:= totalSeg - (totalSeg % 86400)
	totalSeg   	-= Sdias
	Shoras 		:= totalSeg - (totalSeg % 3600)
	totalSeg 	-= Shoras
	Sminutos 	:= totalSeg - (totalSeg % 60)
	totalSeg  	-= Sminutos
	segundos  	:= totalSeg

	dias      	:= Sdias    / 86400
	horas     	:= Shoras   / 3600
	minutos   	:= Sminutos / 60

	segundos 	= totalSeg % 60

	ret = DateTime {
		day : dias,
		hour : horas,
		minutes : minutos,
		seconds : segundos,
		allInSec : (horas * 3600) + (minutos * 60) + (segundos),
	}

	return
}

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