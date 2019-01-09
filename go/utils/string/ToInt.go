package string

import "strconv"
import "fmt"

func (str String) ToInt() (val int) {
    var err error
    val , err = strconv.Atoi(string(str))

    if err != nil {
        fmt.Println(err)
    }

    return
}