package main

type IIterable interface {
    GetNext() IIterable
    GetPrev() IIterable

    SetNext(it *IIterable)
    SetPrev(it *IIterable)

    GetValue() interface{}
    SetValue(it interface{})
}
