package main

type IIterable interface {

    GetNext() IIterable
    GetPrev() IIterable

    SetNext(node IIterable)
    SetPrev(node IIterable)

    GetValue() interface{}
    SetValue(node interface{})

    ToString() string
}
