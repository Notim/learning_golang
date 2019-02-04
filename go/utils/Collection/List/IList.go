package List

type IList interface {
    OrderBy(Action func(Model interface{}) interface{}) (temp List)
    Map(Action func(current interface{}) interface{}) (temp List)
    Lenght() int
    Info() (inf string)
    Foreach(Action func(current interface{}))
    Filter(predicate func(current interface{}) bool) (tempList List)
    Capacity() int
    Any(predicate func(current interface{}) bool) bool
    Add(model ...interface{})
}