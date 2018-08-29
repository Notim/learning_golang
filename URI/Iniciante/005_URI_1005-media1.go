package main
import("fmt")

func main()  {
     var  notaA, notaB float64

     fmt.Scanf("%f\n", &notaA)
     fmt.Scanf("%f\n", &notaB)

     resultado := ((notaA * 3.5) + (notaB * 7.5)) / 11

     fmt.Printf("MEDIA = %.5f\n", resultado)
}
