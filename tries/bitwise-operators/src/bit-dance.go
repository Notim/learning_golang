package main

import (
    "fmt"
    "time"
)

func main()  {
    var val     uint64 = 0xFFFFFFFFFFFFFFFF
    var s         byte = 0x00
    var i         byte = 0x00

    for{
        s++
        for  i = 0x00; i < s; i++{
            fmt.Printf("%2d %064b %064b 0x%016X  %020d\n",s, val, ^val, val, val)
            val = val<<1

            time.Sleep(0x895440)
        }
        for i = 0x00; i < s; i++{

            fmt.Printf("%2d %064b %064b 0x%016X  %020d\n",s,val, ^val, val,val)

            val = val>>1

            time.Sleep(0x895440)
        }

        if s == 0x3f {
            val = ^val // 0xFFFFFFFFFFFFFFFF
            s = 0x00
        }
    }
}
