package main

import "fmt"

type item struct {
    value float64
    next *item
}

func main() {
    var no1 item
    var no2 item
    var no3 item

    no1.value = 5
    no1.next = &no2

    no2.value = 7
    no2.next = &no3

    no3.value = 9
    no3.next = nil

    fmt.Print(no1, no2, no3)
}