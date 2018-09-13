package main
     import "fmt"

     func main(){
          var vector [2]float64
          var quadrante string

          for i := 0; i < len(vector); i++ {
               fmt.Scanf("%f",&vector[i])
          }
          if (vector[0] == 0) && (vector[1] == 0) {
               quadrante = "Origem"
          }else if  (vector[0] == 0) {
               quadrante = "Eixo Y"
          }else if (vector[0] < 0) {
               if vector[1] < 0 {
                    quadrante = "Q3"
               }else if (vector[1] > 0) {
                    quadrante = "Q2"
               } else if (vector[1] == 0){
                    quadrante = "Eixo X"
               }
          }else if (vector[0] > 0) {
               if vector[1] < 0 {
                    quadrante = "Q4"
               }else if (vector[1] > 0) {
                    quadrante = "Q1"
               }else if (vector[1] == 0){
                    quadrante = "Eixo X"
               }
          }
          fmt.Print(quadrante,"\n")
     }
