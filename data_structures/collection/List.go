package main

import "fmt"

type List struct {
    _start *node
    _end   *node
}

/*func (list *List) Append(it interface{}) {
    newItem := &node{value:it}

    if list._start == nil {
        list._start = newItem
        list._end   = newItem

        return
    }

    GetLast(list._start).next = newItem
    list._end = GetLast(list._start)
}*/

func (list *List) Add(item interface{}){
    newNode := &node{ value: item }

    if list._start == nil {
        list._start = newNode
    }

    before := list._end

    if list._end != nil{
        before = list._end
        list._end.next = newNode
    }

    list._end = newNode
    list._end.prev = before
}

func (list *List) ToString() string {
    str := ""

    current := list._start

    for current != nil {
        str += fmt.Sprintf("%v, ", current.value)
        current = current.next
    }

    return "[ " + str + " ]"
}

func (list *List) Length() int {
    if list._start == nil {
        return 0
    }
    return 0
}

func (list *List) GetPrev(it interface{}) interface{} {
    return nil
}

func (list *List) GetNext(it interface{}) interface{} {
    return nil
}