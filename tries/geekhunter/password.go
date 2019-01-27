package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

var reader *bufio.Reader
var (
    line string
    size int
    value map[string]int
)

func init() {
    reader = bufio.NewReader(os.Stdin)
}
func main() {
    loop , _ := strconv.Atoi(ReadLine())

    for i := 0; i < loop; i++ {
        value = make(map[string]int)

        line = ReadLine()

        rep := strings.Split(line, " ")
        size, _ = strconv.Atoi(rep[0])

        word := rep[1]
        limit := (len(word) - size)

        for index := 0; index <= limit; index++  {
            value[string(word[index : index + size])] += 1
        }

        count, hash := 0, ""
        for ha, co :=  range(value) {
            if co > count {
                hash, count = ha, co
            }
        }
        fmt.Println(hash)
   }
}

func ReadLine(msg ...string) string {
    if len(msg) > 0{
        for _, e := range(msg){
            fmt.Print(e)
        }
    }

    input, _ := reader.ReadString('\n')

    return ClearStr(input)
}
func ClearStr(str string) string {
    str = strings.TrimSuffix(str, "\n")
    str = strings.TrimSuffix(str, "\r")
    str = strings.TrimSuffix(str, "\r\n")

    return str
}