package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "strings"

    "github.com/mitchellh/mapstructure"
    "utils/Enumerables/List"

    "./enc"
    "./model"
)


func main() {
    var listPersons List.List = []interface{}{}

    data, err := ioutil.ReadFile("./data/DATABASE") // return []bytes
    if err != nil {
        fmt.Print(err)
    }

    data = enc.Base64Decode(data)

    er := json.Unmarshal(data, &listPersons) // covert to list of maps
    if er != nil {
        panic(er)
    }

    nl := listPersons.Map(func(current interface{}) interface{} {
        var p model.Person = model.Person{}
        if er := mapstructure.Decode(current, &p); er != nil {
            panic(er)
        }

        return p
    })

    (func(log string) {
        fmt.Println(log)
        nl = nl.Filter(func(x interface{}) bool {
            return x.(model.Person).Money > 0 &&
                   strings.Contains(x.(model.Person).Name, "Paulino")
        })
    })("Running")

    (func(log string) {
        fmt.Println(log)
        nl = nl.OrderBy(func(Model interface{}) interface{} {
            return Model.(model.Person).Name
        })
    })("Running")

    nl.Foreach(func(c interface{}) {
        fmt.Printf(
            "#%5d\t| $%6.2f\t| %s\t| %s\n",
            c.(model.Person).ID,
            c.(model.Person).Money,
            Abreviate(c.(model.Person).Name, 17),
            c.(model.Person).Country,
        )
    })
    fmt.Println(nl.Lenght(), "Rows afected of", listPersons.Lenght())
}
func Abreviate(str string, rang int) string {
    if len(str) > rang {
        return str[:rang] + ".."
    }
    return str
}
