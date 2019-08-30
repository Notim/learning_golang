package main

import "time"

// nao fazer desse modo pois pode ocorrer protecao de escrita do objeto enviado para a closure e nao permitir escrita
func ClosureMode(){

    var routine = make(map[string]string)

    go (func(arg string) {
        Infinityloop(arg)
        routine["GO1"] = "Done"
    })("Go 1")

    go (func(arg string) {
        Infinityloop(arg)
        routine["GO2"] = "Done"
    })("Go 2")

    go (func(arg string) {
        Infinityloop(arg)
        routine["GO3"] = "Done"
    })("Go 3")

    go (func(arg string) {
        Infinityloop(arg)
        routine["GO4"] = "Done"
    })("Go 4")

    for {
        if (routine["GO1"] == "Done" &&
            routine["GO2"] == "Done" &&
            routine["GO3"] == "Done" &&
            routine["GO4"] == "Done") {

            break
        }else {
            time.Sleep(1000)
        }
    }
}