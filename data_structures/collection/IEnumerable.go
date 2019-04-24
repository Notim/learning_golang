package main

type IEnumerable interface {
    GetNext(it interface{}) interface{}
    GetPrev(it interface{}) interface{}
    Length() int
}