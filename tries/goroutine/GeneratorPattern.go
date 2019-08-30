package main

import "time"

// padrao goroutine generete
func QueueGeneratorMode(qtdLoop int) <- chan string {
    queue := make(chan string, 3)

    for index := 1; index <= qtdLoop; index++  {
        go func(chan string){
            sec := index * int(time.Second)
            time.Sleep(time.Duration(sec))
            queue <- Infinityloop("GO" + string(qtdLoop))
        }(queue)
    }

    return queue
}