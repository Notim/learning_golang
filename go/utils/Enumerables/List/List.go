package List

type List interface {
    Add(value interface{})
    Filter(predicate func(current interface{}) bool)
    Foreach(action func(element interface{}))
    Lenght() int
}

func (T *ListInt32) Add(value int32) {
    *T = append(*T, value)
}
func (L *ListInt32)  Filter(predicate func(current interface{}) bool){

}
func (L *ListInt32) Foreach(action func(element ListInt32)){

}
func (L *ListInt32) Lenght() int {
    return 0
}

type ListInt32 []int32

type ListInt []int

type ListFloat []float64

type ListString []string
