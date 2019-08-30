package main

import "fmt"

func main() {
    // ClosureMode()

    var loop = 5
    queue := QueueGeneratorMode(loop)

    for i := 0; i < loop ; i++ {
        fmt.Println(<-queue)
    }

    /*
    queue := Mix(
        QueueGeneratorMode(3),
        QueueGeneratorMode(3),
    )
    fmt.Println(<-queue)
    fmt.Println(<-queue)
    fmt.Println(<-queue)
    fmt.Println(<-queue)
    */
    // fmt.Println(theMostFaster())

    // queue := Mixing(falar("Joao"), falar("maria"))

    /*
    fmt.Println(<-queue)
    fmt.Println(<-queue)
    fmt.Println(<-queue)
    fmt.Println(<-queue)
    fmt.Println(<-queue)
    fmt.Println(<-queue)
    fmt.Println(<-queue)
    fmt.Println(<-queue)
    */
}

