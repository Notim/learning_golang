package main

import (
"bufio"
"fmt"
"os"
"reflect"
"strconv"
"strings"
)

type Person struct{
    name string
    age int32
}

func main() {
    var loop int
    reader := bufio.NewReader(os.Stdin)

    loop , err := strconv.Atoi(ReadLine(reader))
    Except(err)

    for i := 0; i < loop; i++ {
        line := ReadLine(reader)
        Except(err)

        var nums []string = strings.Split(line, " ")

        Sort(nums)
        Sort([]int{1,2,3})
        Sort([]float64{1.2,2.5,3.7})
        Sort([]Person{{name:"joao", age: 21},{name:"maria", age:22}})

        /*
            for j := 0; j != (nums[1] - nums[0]); j++  {

            }
        */
    }
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

func Sort(sliceIn interface{}){
    // sliceType  := reflect.TypeOf(sliceIn)
    Val := reflect.ValueOf(sliceIn)
    Type  := reflect.TypeOf(sliceIn)
    swap  := reflect.Swapper(sliceIn)
    length   := Val.Len()

    fmt.Println(Val, Type, swap, length)

    if Type.Kind() == reflect.Slice {

        for ind := 1; ind < length ; ind++ {
            in  := Val.Index(ind - 1).Interface()
            out := Val.Index(ind).Interface()

            switch T := in.(type) {
            case int:
                if T > out.(int){
                    in , out = out, in
                }
                break
            case string:
                if T > out.(string){
                    in  , out = out, in
                }
                break

            case float64:
                if T > out.(float64) {
                    in , out = out, in
                }
                break
            }
        }
    } else {
        os.Exit(1)
    }
}
func GetTypeArray(arr interface{}) reflect.Type {
    return reflect.TypeOf(arr).Elem()
}
func SliceToString(values []interface{}) string {
    s := make([]string, len(values)) // Pre-allocate the right size

    for index := range values {
        s[index] = fmt.Sprintf("%v", values[index])
    }
    return strings.Join(s, ",")
}

//Reverse reverses a slice.
func Reverse(slice *[]interface{}) {
    s := reflect.ValueOf(slice)

    if s.Kind() != reflect.Ptr {
        panic("Must be a pointer")
    }

    sliceLen := s.Elem().Len()
    sliceType := *slice
    for left, right := 0, sliceLen-1; left < right; left, right = left+1, right-1 {
        sliceType[left], sliceType[right] = sliceType[right], sliceType[left]
    }

    s.Elem().Set(reflect.ValueOf(sliceType))
}