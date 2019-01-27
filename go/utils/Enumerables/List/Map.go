package List

func (M *List) Map(Action func(current interface{}) interface{}) (temp List) {
    for _, el := range(*M) {
        temp = append(temp, Action(el))
    }

    return
}

