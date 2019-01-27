package List

import (
    "fmt"
    "reflect"

    "github.com/mitchellh/mapstructure"
)

type info struct {
    TypeOf   interface{}
    ValueOf  interface{}
    Lenght   int
    Capacity int
}

type List []interface{}

func (T *List) To(Type interface{}) (nl List){
    nl = T.Map(func(current interface{}) interface{} {
        t   := reflect.TypeOf(Type)
        p   := reflect.New(t)
        fmt.Println(p)

        fmt.Println(reflect.TypeOf(&p))
        if er := mapstructure.Decode(current, p); er != nil {
            panic(er)
        }
        // fmt.Println(current, p)
        return p
    })

    return
}
