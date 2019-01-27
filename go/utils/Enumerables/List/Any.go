package List

func (M *List) Any(predicate func(current interface{}) bool) bool{
    for _, el := range(*M) {
        if predicate(el) {
            return true
        }
    }
    return false
}