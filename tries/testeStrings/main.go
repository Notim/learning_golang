package main

import (
    "fmt"
    "utils/string"
)

func main() {
    var str string.String = "Hello World"

    fmt.Println(str.Reversed())
    fmt.Println(str.Lenght())
    fmt.Println(str.ToFloat64())

    str = "5.2"
    fmt.Println(str.ToFloat64())

    str = "5"
    fmt.Println(str.ToInt())
}