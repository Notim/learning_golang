package main

import "fmt"

type doubleItem struct {
    before *doubleItem
    value  float64
    after  *doubleItem
}

func main() {
    var no1 doubleItem
    var no2 doubleItem
    var no3 doubleItem
    var no4 = doubleItem {
        value:5,
        before:nil,
        after:&no2,
    }

    no1.value = 5
    no1.after = &no2

    no2.before = &no1
    no2.value  = 7
    no2.after  = &no3

    no3.before = &no2
    no3.value = 9
    no3.after = nil

    fmt.Println(no1, no2, no3, no4)
}