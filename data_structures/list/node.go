package main

import "fmt"

type node struct {
    prev   *node
    value  interface{}
    next   *node
}

func GetLast(item *node) (last *node) {
    if item.next != nil {
        return GetLast(item.next)
    }
    return item
}

func Running(item *node) (retur string) {
    retur = retur + fmt.Sprintf("%v", item.value) + ", "

    if item.next != nil {
        retur = retur + Running(item.next)

        return retur
    }
    return fmt.Sprintf("%v", item.value)
}