
package main
import (
    "fmt"
)
func main() {
    text := "Golang é "
    ou1, out2 := showGolangEhTopper(text)

    fmt.Printf("Hello world %s %s\n", ou1, out2)
}
func showGolangEhTopper(text string) (out1, out2 string) {
    out1 = text  + "topper"
    out2 = "demais"

    return
}
