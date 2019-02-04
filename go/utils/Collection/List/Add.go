package List

func (M *List) Add(model ...interface{}) {
    for _, element := range (model) {
        *M = append(*M, element)
    }
}