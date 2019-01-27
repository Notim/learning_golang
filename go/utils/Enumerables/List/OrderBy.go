package List

import (
    "fmt"
    "sort"
)

func(M *List) OrderBy(Action func(Model interface{}) interface{}) (temp List){
    var mapModel = make(map[interface{}]interface{})
    var key, val interface{}

    for _, el := range(*M) {
        key, val = Action(el), el

        mapModel[key] = val
    }

    var newList = make([]interface{}, 0)

    switch key.(type) {
        case string:{
            var keys []string
            for ha, _ := range(mapModel){
                keys = append(keys, ha.(string))
            }
            sort.Strings(keys)

            for _, el := range(keys) {
                newList = append(newList, mapModel[el])
            }
        }
        case int:{
            var keys []int
            for ha, _ := range(mapModel){
                keys = append(keys, ha.(int))
            }
            sort.Ints(keys)

            for _, el := range(keys) {
                newList = append(newList, mapModel[el])
            }
        }
        case float64:{
            var keys []float64
            for ha, _ := range(mapModel){
                keys = append(keys, ha.(float64))
            }
            sort.Float64s(keys)

            for _, el := range(keys) {
                newList = append(newList, mapModel[el])
            }
        }
        default:{
            fmt.Printf("Sorry, OrderBy cannot support %T\n", key)
            return make(List,0)
        }
    }

    return newList
}