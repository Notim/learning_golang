package main
     import "fmt"
     import "math"

     const PI = 3.14159
     var A, B, C float64
     /*
          a) a área do triângulo retângulo que tem A por base e C por altura.
          b) a área do círculo de raio C. (pi = 3.14159)
          c) a área do trapézio que tem A e B por bases e C por altura.
          d) a área do quadrado que tem lado B.
          e) a área do retângulo que tem lados A e B.
     */
     func main()  {
          fmt.Scanf("%f %f %f\n", &A, &B, &C)
          fmt.Printf("TRIANGULO: %.3f\n", areaTriangulo(A, C))
          fmt.Printf("CIRCULO: %.3f\n", areaCirculo(C))
          fmt.Printf("TRAPEZIO: %.3f\n", areaTrapezio(A, B, C))
          fmt.Printf("QUADRADO: %.3f\n", areaQuadrado(B))
          fmt.Printf("RETANGULO: %.3f\n", areaRetangulo(A, B))

     }
     func areaCirculo(raio float64) float64{
          area := (math.Pow(raio,2) * PI)

          return area
     }
     func areaQuadrado(lado float64) float64{
          area := math.Pow(lado,2)

          return area
     }
     func areaTriangulo(base, altura float64) float64{
          area := ((base * altura) / 2)

          return area
     }
     func areaRetangulo(base, altura float64) float64{
          area := base * altura
          
          return area
     }
     func areaTrapezio(baseMaior, baseMenor, altura float64) float64{
          area := ((baseMaior + baseMenor) / 2) * altura

          return area
     }
