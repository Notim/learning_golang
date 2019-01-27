package List

func (T *List) Filter(predicate func(current interface{}) bool) (tempList List){
    for _, el := range(*T) {
        if predicate(el) {
            tempList = append(tempList, el)
        }
    }

    return
}