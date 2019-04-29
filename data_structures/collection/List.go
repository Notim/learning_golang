package main

import "fmt"

type List struct {
    _start IIterable
    _end   IIterable
}

func (list *List) Add(item interface{}) {
    var newNode IIterable = &node{ value: item }

    if list._start == nil {
        list._start = newNode
    }

    before := list._end

    if list._end != nil{
        before = list._end
        list._end.SetNext(&newNode)
    }

    list._end = newNode
    list._end.SetPrev(&before)
}

func (list *List) ToString() string {
    str := ""

    current := list._start

    for current != nil {
        str += fmt.Sprintf("%v, ", current.GetValue())
        current = current.GetNext()
    }

    return "[ " + str + " ]"
}

func (list *List) Length() int {
    if list._start == nil {
        return 0
    }
    return 0
}