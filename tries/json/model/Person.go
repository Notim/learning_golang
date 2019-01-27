package model

type Person struct {
    ID            int     `json:"Id"`
    Name          string  `json:"Name"`
    Age           int     `json:"Age"`
    CardNumber    int     `json:"CardNumber"`
    Money         float64 `json:"Money"`
    Birth         string  `json:"Birth"`
    Gender        string  `json:"Gender"`
    Country       string  `json:"Country"`
    MaritalStatus string  `json:"MaritalStatus"`
}

