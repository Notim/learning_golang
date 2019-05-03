package main

import "fmt"

type node struct {
    __ref1 IIterable
    __ref2 IIterable
    __value interface{}
}

func (this *node) GetNext() IIterable {
    return this.__ref1
}
func (this *node) GetPrev() IIterable {
    return this.__ref2
}
func (this *node) SetNext(new IIterable){
    if new != nil {
        this.__ref1 = new
    }
    return
}
func (this *node) SetPrev(new IIterable){
    if new != nil {
        this.__ref2 = new
    }
    return
}
func (this *node) GetValue() interface{} {
    return this.__value
}
func (this *node) SetValue(new interface{}) {
    this.__value = new
}
func (this *node) ToString() (str string) {
    str += "\r\nthis ref: " + fmt.Sprintf("%p", this)
    str += "\r\n\tref prev: " + fmt.Sprintf("%p", this.GetPrev())
    str += "\r\n\tref next: " + fmt.Sprintf("%p", this.GetNext())
    str += "\r\n\tvalue: " + fmt.Sprintf("%v", this.__value)

    return
}