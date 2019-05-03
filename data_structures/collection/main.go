package main

import (
    "fmt"
    "time"
)

type Aluno struct{
    index          int64
    Id             int64
    Nome           string
    Sobrenome      string
    DataNascimento time.Time
}

func main() {
    var listItens IList = new(List)
    listItens.Add(Aluno{
        index:0,
        Id:1,
        Nome:"Joao vitor",
        Sobrenome:"paulino martins",
        DataNascimento: time.Now(),
    })
    listItens.Add(Aluno{
        index:1,
        Id:2,
        Nome:"Maria",
        Sobrenome:"Albuquerque da silva",
        DataNascimento: time.Now(),
    }, Aluno{
        index:2,
        Id:3,
        Nome:"Washington",
        Sobrenome:"Almeida Coelho",
        DataNascimento: time.Now(),
    })
    listItens.Add(Aluno{
        index:3,
        Id:4,
        Nome:"Larissa",
        Sobrenome:"Augusta de souza",
        DataNascimento: time.Now(),
    })

    /*
    fmt.Println(listItens.Length())
    fmt.Println(listItens)
    fmt.Println(listItens.ToString())
    fmt.Println(listItens.Get(0))
    fmt.Println(listItens.Get(2))
    */
    listItens.Remove(1)
    fmt.Println(listItens.Length())
    fmt.Println("--- depois de remover o Washington ---")
    listItens.Remove(1)
    fmt.Println(listItens.Length())
    fmt.Println("--- depois de remover o Washington ---")
    listItens.Remove(0)
    fmt.Println(listItens.Length())
    fmt.Println("--- depois de remover o Washington ---")
    listItens.Remove(1)
    // fmt.Println(listItens.Get(4))
}