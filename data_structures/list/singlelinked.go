package main

import "fmt"

type item struct {
    Val   string
    next  *item
}
type list struct {
    _start *item
    _end   *item
}
func (list *list) add(it string) {
    fmt.Println(list._start, list._end)

    newItem := &item{ Val:it }

    if list._start == nil {
        list._start = newItem
    }
    var before, after *item

    before = list._start
    after  = before.next

    for after != nil{
        swap := before
        before = after
        after  = swap.next
    }
    after = newItem

    list._end = after
}
func (list *list) print() string {
    var before, after *item

    fmt.Println(list._start)

    before = list._start
    after  = before.next

    str := "[ "

    for after != nil{
        str += after.Val + ", "

        swap := before
        before = after
        after  = swap.next
    }

    str += " ]"

    return str
}

func main() {
    var listItens list
    listItens.add("5")
    listItens.add("7")
    listItens.add("8")
    listItens.add("2")
    listItens.add("4")

    fmt.Println(listItens.print())

    /*
    var no1 item
    var no2 item
    var no3 item

    no1.value = 5
    no1.next = &no2

    no2.value = 7
    no2.next = &no3

    no3.value = 9
    no3.next = nil

    fmt.Print(no1, no2, no3)
    */
}