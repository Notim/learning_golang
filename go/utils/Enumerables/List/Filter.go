package List

func (T ListInt) Filter(predicate func(current int) bool) (tempList ListInt){
    for index := 0; index < len(T); index++ {
        if predicate(T[index]) {
            tempList = append(tempList, T[index])
        }
    }

    return
}

func (T ListFloat) Filter(predicate func(current float64) bool) (tempList ListFloat){
    for index := 0; index < len(T); index++ {
        if predicate(T[index]) {
            tempList = append(tempList, T[index])
        }
    }

    return
}

func (T ListString) Filter(predicate func(current string) bool) (tempList []string){
    for index := 0; index < len(T); index++ {
        if predicate(T[index]) {
            tempList = append(tempList, T[index])
        }
    }

    return
}
