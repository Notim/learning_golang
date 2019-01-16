package List


func (T *ListInt) Add(value int) {
    *T = append(*T, value)
}

func (T *ListFloat) Add(value float64) {
    *T = append(*T, value)
}

func (T *ListString) Add(value string) {
    *T = append(*T, value)
}