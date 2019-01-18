package main

import (
    "fmt"
    "time"
    "unsafe"
)

/**
 * Virús que lota a memória e o processamento do computador forçando o travamento do sistema
 * Basicamente são criados varias goroutines que executam a mesma função de loop infinito que faz a concatenação de uma string de forma logarítimica
 * Teste com cuidado saporra !!
 * 2019 @ notim <joao.paulino@yahoo.com.br>
**/

func infinityloop(from string) {
    var run [1e8 + 1e8 + 1e8]int32 // 4bytes * [2e9] => 8000bytes (8kbytes)

    for index , _ := range(run){
        run[index] = 0
        // if (index % 100000 == 0){
            //fmt.Printf("%d", 0)
        //}
    }

    // fmt.Printf("%s : %v\n",from, ShowSize(len(run) * 4))//ShowSize(unsafe.Sizeof(run)))
    // fmt.Printf("\n%s : %v\n",from, ShowSize(unsafe.Sizeof(run)))
    defer fmt.Printf("%v ==> %v\n" ,from, ShowSize(unsafe.Sizeof(run)))
}
func ShowSize(val uintptr) (ret string) {
    // reflect.TypeOf(val)

    ret = fmt.Sprintf("%dB", val)

    if (val > 1e9){
        val := float64(val) / float64(1e9)
        ret = fmt.Sprintf("%.1fGB", val)

        return
    }
    if (val > 1e6){
        val := float64(val) / float64(1e6)
        ret = fmt.Sprintf("%.1fMB", val)

        return
    }
    if (val > 1000){
        val := float64(val) / float64(1000)
        ret = fmt.Sprintf("%.1fKB", val)

        return
    }

    return
}

func main() {
    var rontinas map[string]string = make(map[string]string)

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

    for {
        if (rontinas["GO1"] == "Done" && rontinas["GO2"] == "Done" && rontinas["GO3"] == "Done") {
            break
        }else {
            time.Sleep(1000)
        }
    }
}
