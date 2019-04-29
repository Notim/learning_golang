package main

type node struct {
    prev   *IIterable
    next   *IIterable
    value  interface{}
}

func (node *node) GetPrev() IIterable {
    return *node.prev
}

func (node *node) GetNext() IIterable{
    return *node.next
}

func (node *node) SetPrev(inside *IIterable) {
    node.prev = inside
}

func (node *node) SetNext(inside *IIterable){
    node.next = inside
}

func (node *node) GetValue() interface{}{
    return node.value
}

func (node *node) SetValue(it interface{}) {
    node.value = it
}