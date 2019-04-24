package main

type IList interface {
    IEnumerable

    ToString() string
    Add(item interface{})
}