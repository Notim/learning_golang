package main
     import "fmt"

     func main(){
          var valor int

          fmt.Scanf("%d\n", &valor)
          fmt.Print(valor,"\n")

          for i, valorFinal := 0, valor; i < len(moedas); i++ {
               cont, resto = conversorMoeda(valorFinal, moedas[i])
               fmt.Printf("%d nota(s) de R$ %d,00\n", cont, moedas[i])
               valorFinal = resto
          }
     }
     func conversorMoeda(valorEntrada, moderador int) (contador, resto int){
          resto = (valorEntrada % moderador)
          contador =  (valorEntrada - resto) / moderador

          return
     }
