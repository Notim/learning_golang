package Geometry

import (
    "fmt"
)

// Here’s a basic interface for geometric shapes.
type geometry interface {
    area() float64
    perim() float64
}

// If a variable has an interface type, then we can call methods that are in the named interface. Here’s a generic measure function taking advantage of this to work on any geometry.
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}


/*
func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

    The circle and rect struct types both implement the geometry interface so we can use instances of these structs as arguments to measure.

    measure(r)
    measure(c)
}*/
