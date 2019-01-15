package List

import "reflect"

func From(source interface{}) (src interface{}){
    src := reflect.ValueOf(source)

    return
}

type ListInt []int

type ListFloat []float64

type ListString []string
