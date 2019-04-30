package main

import (
    "fmt"
    "time"
)

type Aluno struct{
    Id             int64
    Nome           string
    Sobrenome      string
    DataNascimento time.Time
}

func main() {
    var listItens IList = new(List)

    listItens.Add(&Aluno{
        Id:1,
        Nome:"Joao vitor",
        Sobrenome:"paulino martins",
        DataNascimento: time.Now(),
    })
    listItens.Add(&Aluno{
        Id:2,
        Nome:"Maria",
        Sobrenome:"Albuquerque da silva",
        DataNascimento: time.Now(),
    })
    listItens.Add(&Aluno{
        Id:3,
        Nome:"Washington",
        Sobrenome:"Almeida Coelho",
        DataNascimento: time.Now(),
    })


    fmt.Println(listItens.Length())
    fmt.Println(listItens)
    fmt.Println(listItens.ToString())
    fmt.Println(listItens.Get(0))
    /*
    var reference1, reference2 *node

    n := &node{__value:3}
    reference1, reference2 = n, n

    fmt.Println(*reference1 == *reference2)
    fmt.Printf("%p %p %p\n", reference1, reference2, n)
    */
}