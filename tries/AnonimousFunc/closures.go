package main

import (
    "fmt"
)

func main()  {
    var Connect = func(args ...byte) {
        ip, port := func(por int, args ...byte) ([]byte, []int) {
            var ip    = make([]byte, 4)
            var port  = make([]int, 0)
            port = append(port, por)

            ip = (func(args ...byte) (temp []byte) {
                for _, e := range(args) {
                    temp = append(temp, e)
                }

                return
            })(args...)

            return ip , port
        }(8080, args...)

        fmt.Printf("New IP: %d.%d.%d.%d:%d\n",ip[0],ip[1],ip[2],ip[3], port[0])
    }

    Connect(192,168,0,1)
}