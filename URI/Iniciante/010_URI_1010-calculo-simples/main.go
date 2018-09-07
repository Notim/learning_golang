package main
     import "fmt"

     var codigo, quantidade  int
     var precoUnitario, valorTotal float64 =  0, 0

     func main()  {
          for i := 0; i < 2 ; i++ {
               fmt.Scanf("%d %d %f\n", &codigo, &quantidade, &precoUnitario)
               valorTotal += float64(quantidade) * precoUnitario
          }
          fmt.Printf("VALOR A PAGAR: R$ %.2f\n", valorTotal)
     }
