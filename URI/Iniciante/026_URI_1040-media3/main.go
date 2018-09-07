package main
     import "fmt"

     func main(){
          var vector [4]float64
          pesos := [4]float64{ 2, 3, 4, 1 }
          var media, divisor float64 = 0, 0
          var entradaFinal float64

          for i := 0; i < len(vector); i++ {
               fmt.Scanf("%f",&vector[i])
               media += vector[i] * pesos[i]
               divisor += pesos[i]
          }
          media = media/divisor
          fmt.Printf("Media: %.1f\n", media)

          if (media < 5) {
               fmt.Println("Aluno reprovado.")
          }else if  (media > 6) {
               fmt.Println("Aluno aprovado.")
          }else {
               fmt.Print("Aluno em exame.")
               fmt.Scanf("\n%f",&entradaFinal)
               fmt.Printf("Nota do exame: %.1f\n", entradaFinal)
               media = (media + entradaFinal) / 2
               if media >= 5 {
                    fmt.Print("Aluno aprovado.\n")
               }else{
                    fmt.Print("Aluno reprovado.\n")
               }
               fmt.Printf("Media final: %.1f\n", media)
          }
     }
