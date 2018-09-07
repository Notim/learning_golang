package main
     import "fmt"

     func main(){
          var vector [3]float64

          for i := 0; i < len(vector); i++ {
               fmt.Scanf("%f", &vector[i])
          }
          //vectorPointers := []int{0,0}
          vector2 := sort(vector[0:])

          for i := 0; i < len(vector2); i++ {
               //fmt.Println(vectorPointers[i])
               fmt.Println("vetor 1",vector[i])
               fmt.Println("vetor 2",vector2[i])
          }
     }

     func sort(vector []float64) []float64 {
          var swap float64
          vectorPointers := vector
          for i:= 0 ; i < len(vectorPointers); i++ {
               for j:= 0 ; j < (len(vectorPointers) - 1); j++ {
                    if vectorPointers[j] > vectorPointers[j + 1] {
                         swap = vectorPointers[j]
                         vectorPointers[j] = vectorPointers[j+1]
                         vectorPointers[j+1] = swap
                    }
               }
          }
          return vectorPointers
     }
