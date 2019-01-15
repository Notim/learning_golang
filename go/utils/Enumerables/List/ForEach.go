package List

func (T *ListInt) Foreach(action func(element int)) {
    for _, element := range(*T) {
        action(element)
    }
}
func (T *ListFloat) Foreach(action func(element float64)) {
    for _, element := range(*T) {
        action(element)
    }
}
func (T *ListString) Foreach(action func(element string)) {
    for _, element := range(*T) {
        action(element)
    }
}
