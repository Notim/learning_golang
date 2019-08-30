package main

import "fmt"

func main()  {
    var count uint64 = 0x00

    goto loop

    loop: {
        count++

        if count < 1e5 {

            fmt.Printf("%d running\n", count)
            goto loop
        }
    }
    fmt.Printf("%d Done\n", count)
}