package main

import "fmt"

func main()  {
    var count uint64 = 0x00

    goto start

    start :{
        count++
        fmt.Printf("%X\tstart app\n", count)
        goto middle
    }
    middle:{
        fmt.Printf("%X\tmiddle app\n", count)
        goto more_middle
    }
    more_middle:{
        fmt.Printf("%X\tmore middle app\n", count)
        goto end
    }
    end:{
        fmt.Printf("%X\tend of app\n", count)
        goto start
    }
}