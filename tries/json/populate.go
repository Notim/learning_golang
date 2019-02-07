package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "math"
    "math/rand"
    "os"

    "./enc"
    "./model"
)
import "utils/Enumerables/List"

var arrayNames    []string
var arraySurNames []string
var name    map[int]string
var surname map[int]string

func init() {
    LoadNames()
    LoadSurNames()
}
func main()  {

    var listPersons List.List

    for index := 0; index < 1e4; index++ {
        var person = model.Person{
            ID:index,
            Age: rand.Intn(60),
            Name: GenerateNames(),
            Money : math.Round(rand.Float64() * float64(rand.Intn(150000))),
            Birth : "null",
            CardNumber: rand.Intn(9999),
            MaritalStatus: GenerateStatus(),
            Country: GenerateCountry(),
            Gender: "null",
        }
        listPersons.Add(person)
    }
    j, _ := json.Marshal(listPersons)

    fmt.Printf("%s", enc.Base64Encode(j))
}
func LoadNames(){
    names, err := os.Open("./data/names")
    if err != nil {
        panic(err)
    }

    defer names.Close()

    var lines []string
    scanner := bufio.NewScanner(names)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    arrayNames = make([]string,0)
    for _, name := range (lines){
        arrayNames = append(arrayNames, name)
    }

    name = map[int]string{}

    for in, na := range(arrayNames) {
        name[in] = na
    }
}

func LoadSurNames() {
    names, err := os.Open("./data/surnames")
    if err != nil {
        panic(err)
    }

    defer names.Close()

    var lines []string
    scanner := bufio.NewScanner(names)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    arraySurNames = make([]string,0)
    for _, name := range(lines){

        arraySurNames = append(arraySurNames, name)
    }

    surname = map[int]string{}

    for in, na := range(arraySurNames) {
        surname[in] = na
    }
}
func GenerateNames() string{
    var str = ""
    str += name[rand.Intn(len(name))]
    for i:=0; i < 2; i++ {
        str += " " + surname[rand.Intn(len(surname))]
    }

    return str
}
func GenerateStatus() string{
    var status = map[int]string{}
    status[0] = "Single"
    status[1] = "Married"
    status[2] = "Widow"
    status[3] = "divorced"

    return status[rand.Intn(4)]
}
func GenerateCountry() string{
    var Country = map[int]string{}
    Country[0] = "United Kingdom"
    Country[1] = "United States"
    Country[2] = "Brazil"
    Country[3] = "French"
    Country[4] = "Canada"
    Country[3] = "Germany"
    Country[5] = "French"
    Country[6] = "Italy"
    Country[7] = "Russia"
    Country[8] = "Portugal"
    Country[9] = "Australia"

    return Country[rand.Intn(10)]
}
