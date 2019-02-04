package List

func (M *List) Foreach(Action func(current interface{})){
    for _, el :=  range(*M)  {
        Action(el)
    }
}