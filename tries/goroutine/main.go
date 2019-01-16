package main

import (
    "fmt"
)

func infinityloop(from string) {
    for i := 1; i != 50000; i++ {
        fmt.Println(from, ":", i)
    }
}

func main() {

    go func(arg string) {
        infinityloop(arg)
    }("1 goroutine")

    go func(arg string) {
        infinityloop(arg)
    }("2 goroutine")

    go func(arg string) {
        infinityloop(arg)
    }("3 goroutine")

    go func(arg string) {
        infinityloop(arg)
    }("4 goroutine")

    go func(arg string) {
        infinityloop(arg)
    }("5 goroutine")

    go func(arg string) {
        infinityloop(arg)
    }("6 goroutine")

    go func(arg string) {
        infinityloop(arg)
    }("7 goroutine")

    go func(arg string) {
        infinityloop(arg)
    }("8 goroutine")

    fmt.Scanln()
    fmt.Println("done")
}
