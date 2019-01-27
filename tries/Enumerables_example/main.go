package main

import (
    "encoding/json"
    "fmt"
    "strings"
    "time"

    "utils/Enumerables/List"
)
type person struct {
    Id          int
    Name        string
    Age         int
    CardNumber  int
    Money       float64
    Birth       time.Time
}
func main()  {
    var listPersons List.List

    // Adding persons to List
    listPersons.Add(person {
        Id: 1,
        Name:"Jonh",
        Age:21,
        CardNumber:524458,
        Money:32.50,
        Birth: time.Date(1998,01,12,0,0,0,0,time.UTC),
    }, person {
        Id: 2,
        Name:"mary",
        Age:23,
        CardNumber:522558,
        Money:225.50,
        Birth: time.Date(1997,07,24,0,0,0,0,time.UTC),
    }, person {
        Id: 3,
        Name:"petter",
        Age:32,
        CardNumber:57821,
        Money:65.00,
        Birth: time.Now(),
    }, person {
        Id:4,
        Name:"johnny",
        Age:25,
        CardNumber:57821,
        Money:17.19,
        Birth: time.Now(),
    })

    // filtering by age if is is most than 21
    fmt.Println(listPersons.Filter(func(p interface{}) bool {
        return p.(person).Age > 21
    }))

    // filtering all by name except petter
    fmt.Println(listPersons.Filter(func(m interface{}) bool {
        return m.(person).Name != "petter"
    }))

    // filtering all by name except petter and persons above 30 age
    fmt.Println(listPersons.Filter(func(m interface{}) bool {
        return m.(person).Name != "petter" && m.(person).Age < 30
    }))

    // verifying if exist someone with id 1 (true)
    fmt.Println(listPersons.Any(func(m interface{}) bool {
        return m.(person).Id == 1
    }))

    // verifying if exist someone with id 10 (false)
    fmt.Println(listPersons.Any(func(m interface{}) bool {
        return m.(person).Id == 10
    }))

    // Print Info
    fmt.Println(listPersons.Info())

    // foreach by all piece executing a void method
    // in this case he print name and birthday
    listPersons.Foreach(func(current interface{}) {
        fmt.Println(current.(person).Name ," => ",current.(person).Birth.Format("01-02-2006"))
    })

    // modify some place in the collection
    changed := listPersons.Map(func(x interface{}) interface{} {
        x = person {
            Id        : x.(person).Id,
            CardNumber: 12345,
            Name      : strings.ToUpper(x.(person).Name),
            Age       : x.(person).Age,
            Money     : x.(person).Money,
        }
        return x
    })

    // print changed list with foreach
    changed.Foreach(func(current interface{}) {
        fmt.Println(current)
    })

    // print changed list with foreach
    listPersons.Foreach(func(current interface{}) {
        fmt.Println(current)
    })

    // order List By name
    OrderedList := listPersons.OrderBy(func(x interface{}) interface{} {
        return x.(person).Name
    })

    // write ordered list
    OrderedList.Foreach(func(current interface{}) {
        fmt.Println(current)
    })

    j, _ := json.Marshal(OrderedList)
    fmt.Println(string(j))
}