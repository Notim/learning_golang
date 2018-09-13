package main
     import "fmt"

     func main(){
          var valorEntrada, ano, mes, dia ,resto int

          fmt.Scanf("%d\n", &valorEntrada)

          resto = valorEntrada % 365
          ano = (valorEntrada - resto) / 365
          valorEntrada = resto

          resto = valorEntrada % 30
          mes = (valorEntrada - resto) / 30
          dia = resto
          //1 ano(s)
          //1 mes(es)
          //5 dia(s)

          fmt.Print(ano," ano(s)\n", mes ," mes(es)\n", dia," dia(s)\n")

     }
