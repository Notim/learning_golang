package ListInt

func (T ListInt) Filter(predicate func(current int) bool) (tempList ListInt){
    for index := 0; index < len(T); index++ {
        if predicate(T[index]) {
            tempList = append(tempList, T[index])
        }
    }

    return
}