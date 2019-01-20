package main

import (
    "fmt"
    "time"
    "unsafe"
)

/**
 * Basicamente são criados varias goroutines que executam a mesma função de loop que faz o incremento de um slice de bytes
 * portanto força a reserva de memória do computador forçando o travamento do sistema se estourar a memoria
 * Teste com cuidado saporra !!
 * 2019 @ notim <joao.paulino@yahoo.com.br>
**/
func infinityloop(from string) {
    var run [1e9]int8 // 1byte * 1e9 => 1000000bytes (1GB)

    for index , _ := range(run){
        run[index] = 0
    }
    defer fmt.Printf("%v ==> %v\n" ,from, ShowSize(unsafe.Sizeof(run)))
}
func ShowSize(val uintptr) (ret string) {
    ret = fmt.Sprintf("%dB", val)

    if (val >= 1e9){
        val := float64(val) / float64(1e9)
        ret = fmt.Sprintf("%.1fGB", val)

        return
    }
    if (val >= 1e6){
        val := float64(val) / float64(1e6)
        ret = fmt.Sprintf("%.1fMB", val)

        return
    }
    if (val >= 1000){
        val := float64(val) / float64(1000)
        ret = fmt.Sprintf("%.1fKB", val)

        return
    }

    return
}

func main() {
    var rontinas = make(map[string]string)

    go (func(arg string) {
        infinityloop(arg)
        rontinas["GO1"] = "Done"
    })("Go 1")

    go (func(arg string) {
        infinityloop(arg)
        rontinas["GO2"] = "Done"
    })("Go 2")

    go (func(arg string) {
        infinityloop(arg)
        rontinas["GO3"] = "Done"
    })("Go 3")

    go (func(arg string) {
        infinityloop(arg)
        rontinas["GO4"] = "Done"
    })("Go 4")

    for {
        if (rontinas["GO1"] == "Done" && rontinas["GO2"] == "Done" && rontinas["GO3"] == "Done" && rontinas["GO4"] == "Done") {
            break
        }else {
            time.Sleep(1000)
        }
    }
}
