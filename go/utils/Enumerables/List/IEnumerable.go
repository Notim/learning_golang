package List

type IEnumerable interface {
    Lenght() int
    Capacity() int
    Info() (inf string)

    Add(model ...interface{})
    Filter(Predicate func(Model interface{}) bool)  (tempList []interface{})
    Map(Action func(Model interface{}) interface{}) (temp List)
    Any(Predicate func(Model interface{}) bool) bool

    Foreach(Action func(Model interface{}))
    OrderBy(Action func(Model interface{}) interface{}) (temp List)
}