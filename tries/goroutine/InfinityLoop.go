package main

import (
    "fmt"
    "unsafe"
)

/**
 * Basicamente são criados varias goroutines que executam a mesma função de loop que faz o incremento de um slice de bytes
 * portanto força a reserva de memória do computador forçando o travamento do sistema se estourar a memoria
 * Teste com cuidado saporra !!
 * 2019 @ notim <joao.paulino@yahoo.com.br>
**/
func Infinityloop(from string) string {
    var run [100000000]int8 // 1byte * 1e9 => 1000000bytes (1GB)

    for index , _ := range(run){
        run[index] = 0
    }
    return fmt.Sprintf("%v ==> %v" ,from, ShowSize(unsafe.Sizeof(run)))
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