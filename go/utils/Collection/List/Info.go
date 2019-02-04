package List

import (
    "fmt"
    "reflect"
)
type info struct {
    TypeOf   interface{}
    ValueOf  interface{}
    Lenght   int
    Capacity int
}
func (M *List) Info() (inf string) {
    i := info {
        TypeOf: reflect.TypeOf(M),
        ValueOf: reflect.ValueOf(M),
        Capacity : M.Capacity(),
        Lenght   : M.Lenght(),
    }

    inf =  fmt.Sprintf("Info:\n")
    inf += fmt.Sprintf("==> Len:%v\n", i.Lenght)
    inf += fmt.Sprintf("==> Cap:%v\n", i.Capacity)
    inf += fmt.Sprintf("==> ValueOf:%v\n", i.ValueOf)
    inf += fmt.Sprintf("==> TypeOf:%v\n", i.TypeOf)

    return
}