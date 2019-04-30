package main

import "fmt"

type List struct {
    _start IIterable
    _end   IIterable
}

func (list *List) Add(item interface{}) {
    var newNode IIterable = new(node)
    newNode.SetValue(item)

    if list._start == nil {
        list._start = newNode
    }

    before := list._end

    if list._end != nil {
        before = list._end
        list._end.SetNext(newNode)
    }

    list._end = newNode
    list._end.SetPrev(before)
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

func (list *List) Length() (cont int){
    cont = 0

    current := list._start

    for current != nil {
        cont ++
        current = current.GetNext()
    }
    return
}

func (list *List) Remove(index uint64){

}
func (list *List) Get(index uint64) interface{}{
    if list.Length() == 0 || int(index ) <= list.Length() {
        panic("index out of range!")
    }

    current, value := list._start, list._start.GetValue()

    for in := 0 ; current != nil; in ++{

        current = current.GetNext()
        value   = current.GetValue()

        if in == int(index) {
            return value
        }
    }

    return value
}