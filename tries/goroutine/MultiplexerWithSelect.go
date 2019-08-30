package main

import (
    "fmt"
    "time"
)

func falar(pessoa string) <- chan string {
    chanel := make(chan string)

    go func(chan string) <- chan string {
        for i := 0; i <= 3 ; i++ {
            time.Sleep(time.Second)
            chanel <- fmt.Sprintf("%s falando: %d", pessoa, i)
        }
        return chanel
    }(chanel)

    return chanel
}

func Mixing(inside1, inside2 <- chan string) <- chan string {
    chanel := make(chan string)

    go func(){
        for {
            select {
                case s := <- inside1: chanel <- s
                case s := <- inside2: chanel <- s
            }
        }
    }()
    return chanel
}