package main
     import "fmt"

     func main(){
          var entrada float64
          fmt.Scanf("%f", &entrada)

          if ((entrada >= 0) && (entrada <= 25)){
               fmt.Println("Intervalo [0,25]")
          }else if ((entrada > 25) && (entrada <= 50)){
               fmt.Println("Intervalo (25,50]")
          }else if  ((entrada > 50) && (entrada <= 75)){
               fmt.Println("Intervalo (50,75]")
          }else if ((entrada > 75) && (entrada <= 100)){
               fmt.Println("Intervalo (75,100]")
          } else {
               fmt.Println("Fora de intervalo")
          }
     }
