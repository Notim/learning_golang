package main

import "fmt"

func main()  {
    fmt.Println("it is a print line ftm.Println()")

    v := 5
    fmt.Printf("it is a print line with a formatted type ftm.Printf(%d, v)\n", v)

    var in1 int64

    _, err := fmt.Scanf("%d", &in1)

    if ( err != nil ) {
        fmt.Println(err)
    }

    fmt.Printf("Value %d\n", in1)
}