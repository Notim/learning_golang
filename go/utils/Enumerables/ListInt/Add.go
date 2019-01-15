package ListInt

func (T *ListInt) Add(value int) {
    *T = append(*T, value)
}