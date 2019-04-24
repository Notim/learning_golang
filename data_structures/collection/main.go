package main

import "fmt"

func main() {
    var listItens IList

    listItens = new(List)


    listItens.GetNext(5)
    /*
    listItens.Append(5)
    listItens.Append(7)
    listItens.Append(18)
    listItens.Append(2.2)
    listItens.Append(4)
    listItens.Append("A")
    listItens.Append(listItens)
    listItens.Append(&node{ value : 5 })
    */
    listItens.Add(5)
    listItens.Add(7)
    listItens.Add(18)
    listItens.Add(2.2)
    listItens.Add(4)
    listItens.Add("A")
    listItens.Add(nil)
    listItens.Add(listItens)
    listItens.Add(&listItens)
    listItens.Add(
        &node{
            prev:&node{
                value:nil,
            },
            value:2,
            next: &node{
                value:nil,
            },
        })
    listItens.Add(new(node))

    fmt.Println(listItens.Length())
    fmt.Println(listItens.ToString())

    var reference1, reference2 *node

    n := &node{value:3}
    reference1, reference2 = n, n

    fmt.Println(*reference1 == *reference2)
    fmt.Printf("%p %p %p\n", reference1, reference2, n)
}

func test(action func() string) string {
    return action() + " Foi"
}