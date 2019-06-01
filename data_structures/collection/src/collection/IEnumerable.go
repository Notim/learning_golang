package collection

type IEnumerable interface {
    Length() int
    ToString() string
    Iterate() (uint64, interface{})
}