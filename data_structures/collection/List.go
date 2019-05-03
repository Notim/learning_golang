package main

import "fmt"

type List struct {
    start   IIterable
    end     IIterable
    current IIterable
    value   interface{}
}

func (list *List) Add(items ...interface{}) {
    for _, item := range items {
        var newNode IIterable = new(node)
        newNode.SetValue(item)
        if list.start == nil {
            list.start = newNode
        }
        before := list.end
        if list.end != nil {
            before = list.end
            list.end.SetNext(newNode)
        }
        list.end = newNode
        list.end.SetPrev(before)
    }
}

func (this *List) Remove(index uint64) {
    if this.Length() == 0 || int(index) >= this.Length() {
        panic("index out of range")
    }

    in := 0
    for in, this.current = 0, this.start ; this.current != nil; in ++ {
        if in == int(index) {
            prev        := this.current.GetPrev()
            this.current = this.current.GetNext()

            this.current.SetPrev(prev)
            prev.SetNext(this.current)
        }
        this.current = this.current.GetNext()
        this.end = this.current
    }

    for current := this.start; current != nil ; {
        fmt.Printf("%v\n", current.ToString())

        current = current.GetNext()
    }

    return
}

func (this *List) ToString() string {
    str := ""
    this.current = this.start
    for this.current != nil {
        str += fmt.Sprintf("%v, ", this.current.GetValue())
        this.current = this.current.GetNext()
    }
    return "[ " + str + " ]"
}

func (this *List) Length() (cont int) {
    cont = 0
    this.current = this.start
    for this.current != nil {
        cont ++
        this.current = this.current.GetNext()
    }

    return
}

func (this *List) Iterate() (uint64, interface{}){
    return 0, nil
}

func (this *List) Get(index uint64) interface{} {
    if this.Length() == 0 || int(index) >= this.Length() {
        return nil
    }
    this.current , this.value = this.start, this.start.GetValue()
    for in := 0 ; this.current != nil; in ++ {
        if in == int(index) {
            return this.value
        }
        this.current = this.current.GetNext()
        this.value   = this.current.GetValue()
    }
    return this.value
}