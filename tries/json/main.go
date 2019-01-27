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
                   strings.Contains(x.(model.Person).Country, "United States")
        })
    })("Running")

    (func(log string) {
        fmt.Println(log)
        nl = nl.OrderBy(func(Model interface{}) interface{} {
            return Model.(model.Person).Name
        })
    })("Running")



    nl.Foreach(func(c interface{}) {
        fmt.Printf("$%6.2f => %s\n", c.(model.Person).Money, c.(model.Person).Name)
    })
    fmt.Println(nl.Lenght(), "Rows afected")

}