package List

func (M *List) Capacity() int {
    return cap(*M)
}