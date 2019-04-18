package main

import "fmt"

type node struct {
    value  string
    next   *node
}

type list struct {
    _start *node
    _end   *node
}

func GetLast(item *node) (last *node) {
    if item.next != nil {
        return GetLast(item.next)
    }
    return item
}

func (list *list) Append(it string) {
    newItem := &node{ value: it }

    if list._start == nil {
        list._start = newItem
        list._end   = newItem

        return
    }

    GetLast(list._start).next = newItem
    list._end = GetLast(list._start)
}

func GetNext(item *node) (retur string) {
    retur = retur + item.value

    if item.next != nil {
        retur = retur + GetNext(item.next)

        return retur
    }

    return item.value
}

func (list *list) Print() string {
    fmt.Println(GetLast(list._start))

    return GetNext(list._start)
}

func main() {
    var listItens list

    listItens.Append("5")
    listItens.Append("7")
    listItens.Append("8")
    listItens.Append("2")
    listItens.Append("4")

    fmt.Println(listItens.Print())

    /*
    var node1 node
    var node2 node
    var node3 node
    var node4 node
    var node5 node

    node1.value = "ola"
    node1.next = &node2

    node2.value = " Mundo"
    node2.next = &node3

    node3.value = " filha"
    node3.next = &node4

    node4.value = " da putta"
    node4.next = &node5

    node5.value = " Teste"
    node5.next = nil

    fmt.Println(
        node1.value,
        node1.next.value,
        node1.next.next.value,
        node1.next.next.next.value,
        node1.next.next.next.next,
    )
    fmt.Println(GetLast(&node1))
    */
}