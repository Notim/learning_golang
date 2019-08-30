package main

import (
    "time"
)

func theMostFaster() string {
    select {
        case i1 := <- QueueGeneratorMode(1): return i1
        case i2 := <- QueueGeneratorMode(1): return i2
        case i3 := <- QueueGeneratorMode(1): return i3
        case <- time.After(1000 * time.Millisecond): return "nobody ha come by time"
        // default: return "nobody ha come by default"
    }
}