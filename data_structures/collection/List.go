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
    return struct {
        Title string
    }{
        "title",
    }
}