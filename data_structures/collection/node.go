package main

type node struct {
    prev   *node
    value  interface{}
    next   *node
}

/*
func Running(item *node) (retur string) {
    retur = retur + fmt.Sprintf("%v", item.value) + ", "

    if item.next != nil {
        retur = retur + Running(item.next)

        return retur
    }
    return fmt.Sprintf("%v", item.value)
}*/