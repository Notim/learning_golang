package string

import "strconv"
import "fmt"

func (str String) ToFloat64() (val float64) {
    var err error
    val, err = strconv.ParseFloat(string(str), 64)

    if err != nil {
        fmt.Println(err)
    }

    return
}