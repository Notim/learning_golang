package main
     import "fmt"
     import "math"

     const PI = 3.14159
     var raio float64

     func main()  {
          // (4/3) * pi * R3.
          fmt.Scanf("%f\n", &raio)
          area := (((4/3.0) * PI) * math.Pow(raio, 3))
          fmt.Printf("VOLUME = %.3f\n", area)
     }
