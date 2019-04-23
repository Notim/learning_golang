package main

type enumerable interface {
    Append(it interface{})
    ToString() string
}

type singleLinkedList struct {
    _start *node
    _end   *node
}

func (list *singleLinkedList) Append(it interface{}) {
    newItem := &node{ value: it }

    if list._start == nil {
        list._start = newItem
        list._end   = newItem

        return
    }

    GetLast(list._start).next = newItem
    list._end = GetLast(list._start)
}

func (list *singleLinkedList) ToString() string {
    return "[ " + Running(list._start) + " ]"
}