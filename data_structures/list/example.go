package main

import "fmt"

func main() {
    var listItens singleLinkedList

    listItens.Append(5)
    listItens.Append(7)
    listItens.Append(18)
    listItens.Append(2.2)
    listItens.Append(4)
    listItens.Append("A")
    listItens.Append(singleLinkedList{})
    listItens.Append(node{})

    fmt.Println(listItens.ToString())
}