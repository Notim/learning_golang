package main
     import("fmt")

     var (
          codigo  int
          horas, valorHora, valorTotal float64
     )
     func main()  {
          fmt.Scanf("%d\n %f\n %f\n", &codigo, &horas, &valorHora)

          valorTotal = float64(valorHora * horas)

          fmt.Printf("NUMBER = %d\nSALARY = U$ %.2f\n",codigo, valorTotal)
     }
