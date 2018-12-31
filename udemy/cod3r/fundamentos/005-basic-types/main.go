package main

import (
    "fmt"
    "reflect"
)

func main()  {
    // value types
    /*
        bool

        string

        int
        int8
        int16
        int32
        int64

        uint
        uint8
        uint16
        uint32
        uint64
        uintptr

        byte // alias for uint8

        rune // alias for int32 represents a Unicode code point

        float32
        float64

        complex64
        complex128
    */

    var val
    
    fmt.Println("the type of value 32000 is:", reflect.TypeOf(32000))
    fmt.Println("the type of value 'Abc' is:", reflect.TypeOf("Abc"))
    fmt.Println("the type of value 4.5 is:", reflect.TypeOf(4.5))
}