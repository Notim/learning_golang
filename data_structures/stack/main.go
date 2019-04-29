package main

import "fmt"

type IIterable interface {
    GetNext() IIterable
    GetPrev() IIterable

    SetNext(node IIterable)
    SetPrev(node IIterable)

    GetValue() interface{}
    SetValue(node interface{})
}
type node struct {
    __prev IIterable
    __next IIterable
    __value interface{}
}

func (node node) GetNext() IIterable {
    return node.__next
}

func (node node) GetPrev() IIterable {
    return node.__prev
}

func (node node) SetNext(new IIterable){
    node.__next = new
}
func (node node) SetPrev(new IIterable){
    node.__prev = new
}

func (node node) GetValue() interface{} {
    return node.__value
}

func (node node) SetValue(new interface{}) {
    node.__value = new
}


func main() {
    var node1 IIterable = new(node)
    var node2 IIterable = new(node)

    fmt.Println(node1, node2)

    fmt.Println(node1, node2)

}
