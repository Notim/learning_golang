package main
     import "fmt"
     import "os"
     func main(){
          var in [4]int
          for i:= 0; i < 4; i++ {
               fmt.Scan(&in[i])
          }
          if   (in[1] > in[2]) && (in[3] > in[0]) && ((in[2] + in[3]) > (in[0] + in[1])) && (( in[2] > 0) && ( in[3] > 0)) && (in[0] % 2 == 0){
               fmt.Print("Valores aceitos\n")
               os.Exit(0)
          }
          fmt.Print("Valores nao aceitos\n")
     }
