package Geometry

import "math"

type circle struct {
    radius float64
}

//The implementation for circles.
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}
