package main
     import "fmt"
     import "math"
     var x1, y1, x2, y2, XX, YY, distancia float64
     func main()  {
          fmt.Scanf("%f %f\n%f %f", &x1, &y1, &x2, &y2)
          XX = math.Pow((x2 - x1), 2)
          YY = math.Pow((y2 - y1), 2)
          distancia = math.Sqrt(XX + YY)
          fmt.Printf("%.4f\n",distancia)
     }
